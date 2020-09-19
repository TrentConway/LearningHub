// writing a simple API that returns a message 

package main 

import (
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello World!"}`))
}

func NotFound(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusNotFound)
}
