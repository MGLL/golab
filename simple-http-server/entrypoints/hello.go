package entrypoints

import (
	"github.com/mgll/golab/simple-http-server/handler"
	"github.com/mgll/golab/simple-http-server/middleware"
	"net/http"
)

func init() {
	handler.HandleWithMiddleware("/", middleware.Logging(Hello))
	handler.Handle("/2", AltHello)
}

func Hello(r *http.Request) (int, string) {
	return http.StatusOK, "Hello World!"
}

func AltHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is an alternative hello world!"))
}
