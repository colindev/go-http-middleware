package middleware

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
)

type RecoverMiddleware struct {
	io.Writer
}

func (rm *RecoverMiddleware) Wrap(handler http.HandlerFunc) http.HandlerFunc {

	var writer io.Writer
	if rm.Writer != nil {
		writer = rm
	} else {
		writer = os.Stderr
	}

	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// recover any panic
			if rs := recover(); rs != nil {
				trace := debug.Stack()
				msg := fmt.Sprintf("%s\n%s\n", rs, trace)
				if _, err := io.Copy(writer, strings.NewReader(msg)); err != nil {
					fmt.Println(msg)
					fmt.Println(err)
				}
			}
		}()

		handler(w, r)
	}
}
