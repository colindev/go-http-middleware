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

func Examplestacks_WrapHandler() {

	md := &XMiddleware{stacks: []int{}}

	ms := New(md, md, md)
	ms.Add(md)

	handlerFunc := func(w http.ResponseWriter, r *http.Request) {}

	ms.WrapHandler(handlerFunc).ServeHTTP(nil, nil)

	fmt.Println(md.stacks)
	// Output: [1 2 3 4]

}
