package middleware

import (
	"fmt"
	"github.com/mgll/golab/simple-http-server/handler"
	"net/http"
)

func Logging(h handler.Handler) handler.Handler {
	return func(r *http.Request) (int, string) {
		status, body := h(r)
		fmt.Printf("Request handled with status %d\n", status)
		return status, body
	}
}
