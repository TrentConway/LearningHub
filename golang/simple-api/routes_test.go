package main 

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"
)

func TestHomeRoute(t *testing.T){
	r := mux.NewRouter()
	Routes(r)
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/home", nil)
	if err != nil {
		t.Fatal(err)
	}

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK{
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusOK)
	}
	body := `{"message": "Hello World!"}`
	if rr.Body.String() != body{
	        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), body)
	}
}

func TestInvalidRoute(t *testing.T){
	r := mux.NewRouter()
	Routes(r)
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/invalid", nil)
	if err != nil {
		t.Fatal(err)
	}

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound{
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}


