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

func d(w http.ResponseWriter, r *http.Request) {
	// this call here to ParseForm() makes the data from the form submit avaialbe
	// and it is passed to the template below as r.Form
	// and because it's r.Form, if there's url params, it will get those as well
	// if you want just form info, use r.PostForm
	err := r.ParseForm()
	errorHandle(err)
	w.Header().Set("Dog-Key", "yo, dis dog")
	tpl.ExecuteTemplate(w, "dog.gohtml", r.Form)
}

func c(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	errorHandle(err)
	w.Header().Set("Cat-Key", "yo, dis cat")
	tpl.ExecuteTemplate(w, "cat.gohtml", r.Form)
}

func i(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	errorHandle(err)
	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}

func m(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "me.gohtml", "Matt")
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseGlob("templates/*.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(i))
	http.HandleFunc("/me", m)
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)
	http.ListenAndServe(":8080", nil)
}
