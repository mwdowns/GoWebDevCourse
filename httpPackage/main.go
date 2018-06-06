package main

import (
	"html/template"
	"log"
	"net/http"
)

func errorHandle(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// this call here to ParseForm() makes the data from the form submit avaialbe
	// and it is passed to the template below as r.Form
	err := r.ParseForm()
	errorHandle(err)
	w.Header().Set("Dog-Key", "yo, dis dog")
	// switch r.URL.Path {
	// case "/dog":
	// 	tpl.ExecuteTemplate(w, "dog.gohtml", r.Form)
	// case "/cat":
	// 	tpl.ExecuteTemplate(w, "cat.gohtml", r.Form)
	// default:
	// 	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
	// }
	tpl.ExecuteTemplate(w, "dog.gohtml", r.Form)
}

type hotcat int

func (c hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	errorHandle(err)
	w.Header().Set("Cat-Key", "yo, dis cat")
	tpl.ExecuteTemplate(w, "cat.gohtml", r.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseGlob("*.gohtml"))
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()
	mux.Handle("/dog", d)
	mux.Handle("/cat", c)
	http.ListenAndServe(":8080", mux)
}
