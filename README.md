# http middleware

request middleware handler for native http 

[![Go Report Card](https://goreportcard.com/badge/github.com/colindev/go-http-middleware)](https://goreportcard.com/report/github.com/colindev/go-http-middleware)
[![Build Status](https://travis-ci.org/colindev/go-http-middleware.svg?branch=master)](https://travis-ci.org/colindev/go-http-middleware)
5 [![GoDoc](https://godoc.org/github.com/colindev/go-http-middleware?status.svg)](https://godoc.org/github.com/colindev/go-http-middleware)

### Install

```golang
go get -u github.com/colindev/go-http-middleware
```

### Example

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

    mdw := middleware.New(&AccessMiddleware{})

    http.Handle("/", mdw.WrapHandler(func(w http.ResponseWriter, r *http.Request){
        // handler
    }))

    http.ListenAndServe(":8000", nil)
}

```
