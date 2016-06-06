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
		x.n += 1
		x.stacks = append(x.stacks, x.n)
		handler(w, r)
	}
}

func ExampleWrapHandlerFunc() {

	md := &XMiddleware{stacks: []int{}}

	ms := New(md, md, md)
	ms.Add(md)

	handlerFunc := func(w http.ResponseWriter, r *http.Request) {}

	ms.WrapHandler(handlerFunc).ServeHTTP(nil, nil)

	fmt.Println(md.stacks)
	// Output: [1 2 3 4]

}

func ExampleWrapHandler() {
	md := &XMiddleware{stacks: []int{}}

	ms := New(md, md, md)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	ms.WrapHandler(handler).ServeHTTP(nil, nil)

	fmt.Println(md.stacks)
	// Output: [1 2 3]

}
