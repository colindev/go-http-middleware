# http middleware

request middleware handler for native http 

[![Go Report Card](https://goreportcard.com/badge/github.com/colindev/go-http-middleware)](https://goreportcard.com/report/github.com/colindev/go-http-middleware)
[![Build Status](https://travis-ci.org/colindev/go-http-middleware.svg?branch=master)](https://travis-ci.org/colindev/go-http-middleware)
[![GoDoc](https://godoc.org/github.com/colindev/go-http-middleware?status.svg)](https://godoc.org/github.com/colindev/go-http-middleware)

### Install

```golang
go get -u github.com/colindev/go-http-middleware
```

### Example WrapHandlerFunc

```golang

type AccessMiddleware struct {}

func (am *AccessMiddleware) Wrap(handler http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        access := "............."
        fmt.Println(access)
        handler(w, r)
    }
}

func main() {

    wrapper := middleware.New(&AccessMiddleware{})

    http.Handle("/", wrapper.WrapHandler(func(w http.ResponseWriter, r *http.Request){
        // handler
    }))

    http.ListenAndServe(":8000", nil)
}

```

### Example WrapHandler

```golang
// middlerware 同上

type server struct {}

func (server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // do some thing
}

func main() {
    
    wrapper := middleware.New(&AccessMiddleware{})

    http.ListenAndServe(":8000", wrapper.WrapHandler(server{}))

}

```

