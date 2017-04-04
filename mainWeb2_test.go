package mainWeb

import (
	"html/template"
	"testing"
	"net/http"
	"net/http/httptest"

)

func TestMainHandle(t *testing.T) {

	req, err := http.Request("GET", "/handle", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.HandleFunc()
	Handler := t
	handler := http.HandleFunc(TestMainHandle(t *testing.T))


}

