package server

import (
	"io"
	"net/http"
)

func setRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "This is a test message")
	})
}
