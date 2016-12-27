package middleware

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddlewareFunc(t *testing.T) {

	s := New(MiddlewareFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}))

	req := httptest.NewRequest("GET", "http://127.0.0.1", nil)
	recorder := httptest.NewRecorder()

	s.WrapHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Log(r.URL)
	}).ServeHTTP(recorder, req)

	res := recorder.Result()
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	} else if string(b) != "ok" {
		t.Errorf("%s", b)
	}

}
