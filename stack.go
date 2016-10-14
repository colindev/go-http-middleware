package middleware

import "net/http"

// Wrapper wrap http.Handler / http.HandlerFunc
type Wrapper interface {
	Add(...Middleware)
	WrapHandlerFunc(http.HandlerFunc, ...string) http.HandlerFunc
	WrapHandler(http.Handler) http.Handler
}

type stacks []Middleware

// New return stacks instance
func New(ms ...Middleware) Wrapper {
	s := append(stacks{}, ms...)
	return &s
}

func (s *stacks) Add(ms ...Middleware) {
	*s = append(*s, ms...)
}

func (s *stacks) WrapHandlerFunc(fn http.HandlerFunc, methods ...string) http.HandlerFunc {

	for i := len(*s) - 1; i >= 0; i-- {
		fn = (*s)[i].Wrap(fn)
	}

	allowMethods := map[string]bool{}
	for _, m := range methods {
		allowMethods[m] = true
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// 預設不檢查
		if len(allowMethods) > 0 && !allowMethods[r.Method] {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fn(w, r)
	}
}

func (s *stacks) WrapHandler(handler http.Handler) http.Handler {
	return s.WrapHandlerFunc(handler.ServeHTTP)
}
