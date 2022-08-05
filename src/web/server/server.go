package server

import (
	"net/http"
)

func Run() {

	mux := http.NewServeMux()

	setRoutes(mux)

	http.ListenAndServe("0.0.0.0:8086", mux)
}
