package mainWeb


import (

"testing"
"net/http"
"net/http/httptest"

)


func TestformInputHandler(t *testing.T) {


	req, err := http.NewRequest("GET", "/satus", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandleFunc(InputHandler)


	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf (
			status, http.StatusOK)
	}

	expected := "status"
	if rr.Body.String() != expected {
		t.Errorf(
			rr.Body.String(), expected)
	}
}

