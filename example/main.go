package main

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/colindev/go-http-middleware"
)

type AccessMiddleware struct{}

func (am *AccessMiddleware) Wrap(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		access := fmt.Sprintf("%s ", r.URL.String())
		fmt.Println(access)
		handler(w, r)
	}
}

type RecoverMiddleware struct{}

func (rm *RecoverMiddleware) Wrap(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// recover any panic
			if rs := recover(); rs != nil {
				trace := debug.Stack()
				fmt.Printf("%s\n%s\n", rs, trace)
			}
		}()

		handler(w, r)
	}
}

func main() {

	mdw := middleware.New(&AccessMiddleware{}, &RecoverMiddleware{})

	http.Handle("/", mdw.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// handler
		panic("test panic recover")
	}))

	http.ListenAndServe(":8000", nil)
}
