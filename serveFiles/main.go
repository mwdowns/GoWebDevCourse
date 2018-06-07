package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/sloth", sloth)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func sloth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	io.WriteString(w, `
		<img src="/resources/felice.jpg">
	`)
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
