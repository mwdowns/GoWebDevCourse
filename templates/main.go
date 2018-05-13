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

type being struct {
	Name string
	Type string
}

type favorites struct {
	FavoriteFood   string
	FavoritePerson string
}

type items struct {
	Beings    []being
	Favorites []favorites
}

func main() {
	// name := os.Args[1]

	// tpl := fmt.Sprint(`
	// <!DOCTYPE html>
	// <html lang="en">
	// <head>
	// <meta charset="UTF-8">
	// <title>Hello World!</title>
	// </head>
	// <body>
	// 	<h1>` + name + `</h1>
	// </body>
	// </html>
	// `)
	// nf, err := os.Create("index.html")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// defer nf.Close()

	// io.Copy(nf, strings.NewReader(tpl))

	// tpl, err := template.ParseGlob("templates/*.gohtml")
	// errorHandler(err)

	people := []string{"Jeri", "Matt", "Doodins", "Stinkins"}
	people2 := []being{
		{Name: "Jeri", Type: "person"},
		{Name: "Matt", Type: "person"},
		{Name: "Doodins", Type: "cat"},
		{Name: "Stinkins", Type: "cat"},
	}
	favorites := []favorites{
		{FavoriteFood: "Vegan/NonDairy", FavoritePerson: "Doodins"},
		{FavoriteFood: "Goldfish", FavoritePerson: "Stikins"},
		{FavoriteFood: "Milkins", FavoritePerson: "Jeri"},
		{FavoriteFood: "Tunins", FavoritePerson: "Matt"},
	}
	stuff := items{
		Beings:    people2,
		Favorites: favorites,
	}

	nf1, err := os.Create("index.html")
	errorHandler(err)

	nf2, err := os.Create("about.html")
	errorHandler(err)

	nf3, err := os.Create("about2.html")
	errorHandler(err)

	err = tpl.ExecuteTemplate(nf1, "tpl.gohtml", people)
	errorHandler(err)

	err = tpl.ExecuteTemplate(nf2, "tpl2.gohtml", people2)
	errorHandler(err)

	err = tpl.ExecuteTemplate(nf3, "tpl3.gohtml", stuff)
	errorHandler(err)
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
