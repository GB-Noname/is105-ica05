package mainWeb


import (

"testing"
"net/http"
"net/http/httptest"

)


func TestStatusHandle(t *testing.T) {


	req, err := http.NewRequest("GET", "/satus", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandleFunc(StatusHandle)


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

