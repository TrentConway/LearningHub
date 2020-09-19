// writing a simple API that returns a message 

package main

import (
	"log"
	"net/http"
)

// handles the request
func root(w http.ResponseWriter, r *http.Request){
	// header
	w.Header().Set("Content-Type","application/json")
	// status code
	w.WriteHeader(http.StatusOK)
	// http message body
	w.Write([]byte(`{"message": "hello-world"}`))
}


// can access the request on port 8080
func main() {
	// server function
	http.HandleFunc("/", root)
	// defines the port
	log.Fatal(http.ListenAndServe(":8080",nil))
}

// type server struct{}
// root := &server{}
// http.Handle("/", root)
