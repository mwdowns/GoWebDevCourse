package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	panicErr(err)

	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)
	logErr(err)

	fmt.Println(string(bs))
}

func panicErr(e error) {
	if e != nil {
		panic(e)
	}
}

func logErr(e error) {
	if e != nil {
		log.Println(e)
	}
}
