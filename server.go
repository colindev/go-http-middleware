package middleware

import "net/http"

type stacks []Middleware

func New(ms ...Middleware) *stacks {
	s := append(stacks{}, ms...)
	return &s
}

func (s *stacks) HandlerFunc(handler http.HandlerFunc) http.HandlerFunc {
	for i := len(*s) - 1; i >= 0; i-- {
		handler = (*s)[i].Wrap(handler)
	}

	return handler
}

func (s *stacks) Add(ms ...Middleware) {
	*s = append(*s, ms...)
}
