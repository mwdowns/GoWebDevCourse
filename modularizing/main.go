package main

import (
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	// nf, err := os.Create("index.html")
	// errorHandler(err)

	err := tpl.ExecuteTemplate(os.Stdout, "main.gohtml", "Matt")
	errorHandler(err)
	fmt.Println("modularizing!")
}
