package controllers

import (
	// "fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, req *http.Request) {
	// fmt.Fprintf(w, "Welcome to the home page!")
	tmpl, err := template.ParseFiles("src/old.thing/resources/views/index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
