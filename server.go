package middleware

import "net/http"

type stacks []Middleware

func (s stacks) HandlerFunc(handler http.HandlerFunc) http.HandlerFunc {
	for i := len(s) - 1; i >= 0; i-- {
		handler = s[i].Wrap(handler)
	}

	return handler
}

func New(ms ...Middleware) stacks {
	return append(stacks{}, ms...)
}