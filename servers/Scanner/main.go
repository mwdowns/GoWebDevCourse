package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	fmt.Println("yo, servers")
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		// io.WriteString(conn, "\nHello from TCP server\n")
		// fmt.Fprintln(conn, "How is your day?")
		// fmt.Fprintf(conn, "%v", "Well, I hope!")

		// conn.Close()

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("CONN TIMEOUT")
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()

	fmt.Println("Finished")
}