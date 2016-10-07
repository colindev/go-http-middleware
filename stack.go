package middleware

import "net/http"

// Wrapper wrap http.Handler / http.HandlerFunc
type Wrapper interface {
	Add(...Middleware)
	WrapHandlerFunc(http.HandlerFunc) http.HandlerFunc
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

func (s *stacks) WrapHandlerFunc(fn http.HandlerFunc) http.HandlerFunc {

	for i := len(*s) - 1; i >= 0; i-- {
		fn = (*s)[i].Wrap(fn)
	}

	return fn
}

func (s *stacks) WrapHandler(handler http.Handler) http.Handler {

	fn := s.WrapHandlerFunc(handler.ServeHTTP)

	return http.HandlerFunc(fn)
}
