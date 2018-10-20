package controllers

import (
	"html/template"
	"net/http"
)

func HomeIndex(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("resources/views/test-login-mylc.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Static() {
	http.FileServer(http.Dir("static"))
}
