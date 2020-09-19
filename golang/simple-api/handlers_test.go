package main 

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T){ 
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RootHandler)

	handler.ServeHTTP(rr, req)

	// Test Status Code
	if status := rr.Code; status != http.StatusOK{
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Test the payload
	body := `{"message": "Hello World!"}`
	if rr.Body.String() != body{
	        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), body)
	}

}

