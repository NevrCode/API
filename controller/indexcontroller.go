package controller

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	temp, err := template.ParseFiles("index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
