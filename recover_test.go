package middleware

import (
	"bytes"
	"net/http"
	"testing"
)

func TestRecover(t *testing.T) {

	writer := bytes.NewBuffer([]byte{})

	handler := (&RecoverMiddleware{writer}).Wrap(func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	handler(nil, nil)

	if writer.Len() == 0 {
		t.Error("buffer must not be empty")
	}

	line, err := writer.ReadString('\n')
	if err != nil {
		t.Error(err)
		return
	}

	if line != "test\n" {
		t.Errorf("first line must be string(test), but %#v", line)
	}
}
