package main

import (
	"html/template"
	"log"
	"os"
)

type Page struct {
	Title string
}

var templates = template.Must(template.ParseGlob("templates/*"))

func main() {

	p := Page{Title: "Heading"}
	err := templates.ExecuteTemplate(os.Stdout, "index.html", p)
	if err != nil {
		log.Fatal("Cannot Get View ", err)
	}

}