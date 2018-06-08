package main

import (
	"io"
	"net/http"
)

func main() {
	fs := http.StripPrefix("/resources", http.FileServer(http.Dir("./assets")))
	http.HandleFunc("/", index)
	http.HandleFunc("/sloth", sloth)
	http.Handle("/resources/", fs)
	http.ListenAndServe(":8080", nil)
}

func sloth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	io.WriteString(w, `
		<img src="/resources/felice.jpg">
	`)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// func slothPic(w http.ResponseWriter, r *http.Request) {
// 	f, err := os.Open("felice.jpg")
// 	if err != nil {
// 		http.Error(w, "file not found", 404)
// 		return
// 	}
// 	defer f.Close()

// 	io.Copy(w, f)
// }
