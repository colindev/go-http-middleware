package middleware

import "net/http"

type Middleware interface {
	Wrap(http.HandlerFunc) http.HandlerFunc
}

type MiddlewareFunc func(w http.ResponseWriter, r *http.Request)

func (m MiddlewareFunc) Wrap(fn http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		m(w, r)
		fn(w, r)
	}
}
