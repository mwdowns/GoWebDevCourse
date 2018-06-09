package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var tpl *template.Template

// type person struct {
// 	FirstName string
// 	LastName  string
// 	Subscribe bool
// }

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}

func main() {
	// http.HandleFunc("/", v)
	// http.HandleFunc("/", p)
	http.HandleFunc("/", rf)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

// func v(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	io.WriteString(w, `
// 		<form method="post">
// 		<input type="text" name="q">
// 		<input type="submit">
// 		</form><br>
// 		`+v)
// }

// func p(w http.ResponseWriter, r *http.Request) {
// 	f := r.FormValue("first")
// 	l := r.FormValue("last")
// 	s := r.FormValue("subscribe") == "on"

// 	err := tpl.ExecuteTemplate(w, "index_p.gohtml", person{f, l, s})
// 	errorHandler(err)
// }

func rf(w http.ResponseWriter, r *http.Request) {
	var s string
	fmt.Println(r.Method)
	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("file")
		errorHandler(err)
		defer f.Close()

		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr:", err)

		bs, err := ioutil.ReadAll(f)
		errorHandler(err)

		s = string(bs)
	}

	err := tpl.ExecuteTemplate(w, "index_rf.gohtml", s)
	errorHandler(err)
}
