# http middleware

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

    http.Handle("/", mdw.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        // handler
    }))

    http.ListenAndServe(":8000", nil)
}

```
