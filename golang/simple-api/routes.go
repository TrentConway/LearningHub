package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/home", RootHandler).Methods(http.MethodGet)
	r.NotFoundHandler = http.HandlerFunc(NotFound)
}
