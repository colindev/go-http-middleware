package middleware

import (
	"net/http"
	"testing"
)

type TestMiddleware struct {
	history []string
}

func (tmw *TestMiddleware) Wrap(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmw.history = append(tmw.history, r.URL.String())
		handler(w, r)
	}
}

func TestIf(t *testing.T) {

	tmw := &TestMiddleware{make([]string, 0)}

	server := New(&IfMiddleware{
		Condition: func(r *http.Request) bool {

			if r == nil {
				return false
			}

			return r.URL.Path == "/true"
		},

		IfTrue: tmw,

		IfFalse: nil,
	})

	app := http.Handler(server.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))

	app.ServeHTTP(nil, nil)
	req, _ := http.NewRequest("GET", "http://127.0.0.1/true", nil)
	app.ServeHTTP(nil, req)
	req, _ = http.NewRequest("GET", "http://127.0.0.1/x", nil)
	app.ServeHTTP(nil, req)

	if len(tmw.history) != 1 {
		t.Error("only one request can pass", tmw.history)
	}
}
