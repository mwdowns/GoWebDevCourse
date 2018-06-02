package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	errHandle(err)

	defer li.Close()

	for {
		conn, err := li.Accept()
		errHandle(err)
		go handle(conn)
	}

}

func errHandle(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func handle(conn net.Conn) {

	defer conn.Close()

	body := constructBody(request(conn))
	response(conn, body)
}

func request(conn net.Conn) (method string, url string) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		p := strings.Fields(scanner.Text())
		return p[0], p[1]
	}
	return "null", "null"
}

func response(conn net.Conn, body string) {
	// body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title>Hello World</title></head><body><h1>Hello, Matt</h1></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func constructBody(method string, url string) string {
	m := ``
	u := ``

	if method == "GET" {
		m = `<p>GET</p>`
	} else {
		m = `<p>Not a GET method</p>`
	}

	if url == "/" {
		u = `<p>URL is home</p>`
	} else {
		u = `<p>You're not home</p>`
	}

	return `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title>Hello World</title></head><body>` + m + u + `<h1>Hello, Matt</h1></body></html>`
}
