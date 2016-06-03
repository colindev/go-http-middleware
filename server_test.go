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

func Example() {

	md := &XMiddleware{stacks: []int{}}

	server := New(md, md, md)

	http.Handler(server.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})).ServeHTTP(nil, nil)

	fmt.Println(md.stacks)
	// Output: [1 2 3]

}
