package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
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

func (server) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func ExampleWrapper_WrapHandlerFunc() {

	md := &XMiddleware{stacks: []int{}}
	ms := New(md, md, md)
	ms.Add(md)

	recorder := httptest.NewRecorder()
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {}
	ms.WrapHandlerFunc(handlerFunc)(recorder, httptest.NewRequest("GET", "http://127.0.0.1", nil))
	fmt.Println("demo WrapHandlerFunc")
	fmt.Println(md.stacks)

	// Output:
	// demo WrapHandlerFunc
	// [1 2 3 4]

}

func TestWrapper_WrapHandlerFunc(t *testing.T) {
	ms := New()

	recorder := httptest.NewRecorder()
	ms.WrapHandlerFunc(nil, "PUT")(recorder, httptest.NewRequest("GET", "http://127.0.0.1", nil))
	if code := recorder.Result().StatusCode; code != http.StatusBadRequest {
		t.Errorf("expect status code = %d, but got %d", http.StatusBadRequest, code)
	}
}

func ExampleWrapper_WrapHandler() {

	md := &XMiddleware{stacks: []int{}}
	ms := New(md, md, md)

	recorder := httptest.NewRecorder()
	fmt.Println("demo WrapHandler")
	ms.WrapHandler(server{}).ServeHTTP(recorder, httptest.NewRequest("GET", "http://127.0.0.1", nil))
	fmt.Println(md.stacks)

	// Output:
	// demo WrapHandler
	// [1 2 3]
}
