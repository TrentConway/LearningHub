
package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/home", RootHandler).Methods(http.MethodGet)
	r.NotFoundHandler = http.HandlerFunc(NotFound)
}

func main() {
	r := mux.NewRouter()
	Routes(r)
    log.Fatal(http.ListenAndServe(":8080",r))

}

