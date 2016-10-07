package middleware

import (
	"fmt"
	"net/http"
)

type XMiddleware struct {
	stacks []int
	n      int
}

func (x *XMiddleware) Wrap(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		x.n++
		x.stacks = append(x.stacks, x.n)
		handler(w, r)
	}
}

type server struct{}

func (server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[server] ServeHTTP", w, r)
}

func ExampleWrapper_WrapHandlerFunc() {

	md := &XMiddleware{stacks: []int{}}
	ms := New(md, md, md)
	ms.Add(md)

	handlerFunc := func(w http.ResponseWriter, r *http.Request) {}
	ms.WrapHandlerFunc(handlerFunc)(nil, nil)
	fmt.Println("demo WrapHandlerFunc")
	fmt.Println(md.stacks)

	// Output:
	// demo WrapHandlerFunc
	// [1 2 3 4]

}

func ExampleWrapper_WrapHandler() {

	md := &XMiddleware{stacks: []int{}}
	ms := New(md, md, md)

	fmt.Println("demo WrapHandler")
	ms.WrapHandler(server{}).ServeHTTP(nil, nil)
	fmt.Println(md.stacks)

	// Output:
	// demo WrapHandler
	// [server] ServeHTTP <nil> <nil>
	// [1 2 3]
}
