package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

var tpl *template.Template

type person struct {
	FirstName string
	LastName  string
	Subscribe bool
}

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
	http.HandleFunc("/", i)
	http.HandleFunc("/v", v)
	http.HandleFunc("/p", p)
	http.HandleFunc("/rf", rf)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func i(w http.ResponseWriter, r *http.Request) {
	name := "Matt"
	checkCookie(w, r)
	err := tpl.ExecuteTemplate(w, "index_i.gohtml", name)
	errorHandler(err)
}

func v(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("q")
	checkCookie(w, r)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	io.WriteString(w, `
		<a href="/">Home</a>
		<form method="post">
		<input type="text" name="q">
		<input type="submit">
		</form><br>
		`+value)
}

func p(w http.ResponseWriter, r *http.Request) {
	f := r.FormValue("first")
	l := r.FormValue("last")
	s := r.FormValue("subscribe") == "on"
	checkCookie(w, r)
	err := tpl.ExecuteTemplate(w, "index_p.gohtml", person{f, l, s})
	errorHandler(err)
}

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
	checkCookie(w, r)
	err := tpl.ExecuteTemplate(w, "index_rf.gohtml", s)
	errorHandler(err)
}

func checkCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("matt-cookie")
	if err == http.ErrNoCookie {
		http.SetCookie(w, &http.Cookie{
			Name:  "matt-cookie",
			Value: "0",
		})
		return
	}
	oldVal, _ := strconv.Atoi(c.Value)
	c.Value = strconv.Itoa(oldVal + 1)
	http.SetCookie(w, c)
}
