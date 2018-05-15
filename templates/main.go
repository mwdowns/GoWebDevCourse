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

type templates struct {
	pageName     string
	templateName string
}

type being struct {
	Name   string
	Animal string
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

	// people := []string{"Jeri", "Matt", "Doodins", "Stinkins"}
	pages := []templates{
		{pageName: "index.html", templateName: "tpl.gohtml"},
		{pageName: "about.html", templateName: "tpl2.gohtml"},
		{pageName: "about2.html", templateName: "tpl3.gohtml"},
	}
	beings := []being{
		{Name: "Jeri", Animal: "person"},
		{Name: "Matt", Animal: "person"},
		{Name: "Doodins", Animal: "cat"},
		{Name: "Stinkins", Animal: "cat"},
		{Name: "Rovi", Animal: "dog"},
		{Name: "Danny", Animal: "good boy"},
	}
	favorites := []favorites{
		{FavoriteFood: "Vegan/NonDairy", FavoritePerson: "Doodins"},
		{FavoriteFood: "Goldfish", FavoritePerson: "Stikins"},
		{FavoriteFood: "Milkins", FavoritePerson: "Jeri"},
		{FavoriteFood: "Tunins", FavoritePerson: "Matt"},
		{FavoriteFood: "Turkey Neck", FavoritePerson: "Jeri"},
		{FavoriteFood: "Buscuit", FavoritePerson: "Matt"},
	}
	stuff := items{
		Beings:    beings,
		Favorites: favorites,
	}

	for _, files := range pages {
		nf, err := os.Create(files.pageName)
		errorHandler(err)

		err = tpl.ExecuteTemplate(nf, files.templateName, stuff)
		errorHandler(err)
	}

	// nf1, err := os.Create("index.html")
	// errorHandler(err)

	// nf2, err := os.Create("about.html")
	// errorHandler(err)

	// nf3, err := os.Create("about2.html")
	// errorHandler(err)

	// err = tpl.ExecuteTemplate(nf1, "tpl.gohtml", people)
	// errorHandler(err)

	// err = tpl.ExecuteTemplate(nf2, "tpl2.gohtml", people2)
	// errorHandler(err)

	// err = tpl.ExecuteTemplate(nf3, "tpl3.gohtml", stuff)
	// errorHandler(err)
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
