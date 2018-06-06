package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// this call here to ParseForm() makes the data from the form submit avaialbe
	// and it is passed to the template below as r.Form
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Set("Matt-key", "yo, dis Matt")
	switch r.URL.Path {
	case "/dog":
		tpl.ExecuteTemplate(w, "dog.gohtml", r.Form)
	case "/cat":
		tpl.ExecuteTemplate(w, "cat.gohtml", r.Form)
	default:
		tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
	}
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseGlob("*.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
