package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/colindev/go-http-middleware"
)

type accessMiddleware struct{}

func (am *accessMiddleware) Wrap(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		access := fmt.Sprintf("%s ", r.URL.String())
		log.Println(access)
		handler(w, r)
	}
}

func main() {

	var (
		addr    string
		useCors bool
	)
	flag.StringVar(&addr, "addr", ":8000", "http listen address")
	flag.BoolVar(&useCors, "cors", false, "use cors")
	flag.Parse()

	f, err := os.OpenFile("/tmp/error.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	cm := &middleware.CorsMiddleware{
		OriginValidator: func(r *http.Request) (string, bool) {
			return "*", true
		},
	}

	mdw := middleware.New(
		&accessMiddleware{},
		&middleware.RecoverMiddleware{Writer: f},
	)

	if useCors {
		mdw.Add(cm)
	}

	http.HandleFunc("/panic", mdw.WrapHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// handler
		log.Println("test panic")
		panic("test panic recover")
	}))

	http.HandleFunc("/cors", mdw.WrapHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}))

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Println("[http]", err)
	}
}
