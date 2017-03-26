package main

import (
	"html/template"
	"net/http"
	"path"
	"encoding/json"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {
	go http.HandleFunc("/", foo)
	http.ListenAndServe(":8001", nil)
	//http.HandleFunc("/", foo2)
}

func foo(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", "index.html")

	// Note that the layout file must be the first parameter in ParseFiles
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func foo2(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
}