package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", blah)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func blah(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<form method="post">
		<input type="text" name="q">
		<input type="submit">
		</form><br>
		`+v)
}
