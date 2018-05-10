package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]

	tpl := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`)
	nf, err := os.Create("index.html")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))
}
