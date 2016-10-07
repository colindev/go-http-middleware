package main

import (
	"fmt"
	"net/http"
	"os"

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

func main() {

	f, err := os.OpenFile("/tmp/error.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	mdw := middleware.New(&AccessMiddleware{}, &middleware.RecoverMiddleware{f})

	http.Handle("/", mdw.WrapHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// handler
		panic("test panic recover")
	}))

	http.ListenAndServe(":8000", nil)
}
