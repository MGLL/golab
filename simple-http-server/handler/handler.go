package handler

import (
	"net/http"
)

type Handler func(r *http.Request) (int, string)

func Handle(p string, f func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(p, f)
}

func HandleWithMiddleware(p string, h Handler) {
	http.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		status, body := h(r)
		w.WriteHeader(status)
		w.Write([]byte(body))
	})
}
