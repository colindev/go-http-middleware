package middleware

import "net/http"

type IfMiddleware struct {
	Condition func(r *http.Request) bool

	IfTrue Middleware

	IfFalse Middleware
}

func (imw *IfMiddleware) Wrap(handler http.HandlerFunc) http.HandlerFunc {

	if imw.Condition == nil {
		panic("IfMiddleware miss Condition")
	}

	var trueHandler http.HandlerFunc
	if imw.IfTrue != nil {
		trueHandler = imw.IfTrue.Wrap(handler)
	} else {
		trueHandler = handler
	}

	var falseHandler http.HandlerFunc
	if imw.IfFalse != nil {
		falseHandler = imw.IfFalse.Wrap(handler)
	} else {
		falseHandler = handler
	}

	return func(w http.ResponseWriter, r *http.Request) {

		if imw.Condition(r) {

			trueHandler(w, r)

		} else {

			falseHandler(w, r)

		}
	}
}
