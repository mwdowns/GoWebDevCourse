package main

import (
	"fmt"
	"os"
	"text/template"
)

type animal struct {
	Name, Type string
}

type food struct {
	Name string
}

type person struct {
	Name            string
	FavoriteNumber  int
	FavoriteAnimals []animal
	FavoriteFoods   []food
}

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

	p := []person{
		person{
			Name:           "Matt",
			FavoriteNumber: 13,
			FavoriteAnimals: []animal{
				animal{Name: "Stinkins", Type: "cat"},
				animal{Name: "Dandy Dan", Type: "dog"},
			},
			FavoriteFoods: []food{
				food{Name: "Japanese"},
				food{Name: "Indian"},
			},
		},
		person{
			Name:           "Jeri",
			FavoriteNumber: 6,
			FavoriteAnimals: []animal{
				animal{Name: "Doodins", Type: "cat"},
				animal{Name: "Rovi", Type: "dog"},
			},
			FavoriteFoods: []food{
				food{Name: "Vegan/Non-dairy"},
				food{Name: "Thai"},
			},
		},
	}

	nf, err := os.Create("index.html")
	errorHandler(err)

	err = tpl.ExecuteTemplate(nf, "main.gohtml", p)
	errorHandler(err)
	fmt.Println("modularizing!")
}
