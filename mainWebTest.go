package main

import (
	"html/template"
	"net/http"
	"path"
	"encoding/json"
	"fmt"

)

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {
	//http.HandleFunc("/search", search)
	http.HandleFunc("/", foo)
	http.HandleFunc("/search", foo2)
	http.HandleFunc("/AltSubmit", formInputHandler)
	http.HandleFunc("/ttt", foo2)
	http.HandleFunc("/ddd", foo2)
	http.HandleFunc("/fff", foo2)
	http.ListenAndServe(":8001", nil)
	//go http.HandleFunc("/", foo2)
}

func foo(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", "indexTest.html")

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
	r.ParseForm()  //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	name := r.Form.Get("name")
	fmt.Println(name)
	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
}

func formInputHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	for k := range r.Form {
		if k == "List" {
		fmt.Println("testetettetetetetet")
			fmt.Println(k)

	} else if k == "ShowProg" {
			fmt.Println("ABABBABBABABABBA")
			fmt.Println(k)
	} else if k == "ShowCode" {
			fmt.Println("This is code")
			fmt.Println(k)
		}
	}
}

func foo4(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
}

func foo5(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
}
