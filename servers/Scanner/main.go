package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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

		// go handle(conn)
		// go rot13handle(conn)
		go dbhandle(conn)
	}
}

// func handle(conn net.Conn) {
// 	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
// 	if err != nil {
// 		log.Println("CONN TIMEOUT")
// 	}
// 	scanner := bufio.NewScanner(conn)
// 	for scanner.Scan() {
// 		ln := scanner.Text()
// 		fmt.Println(ln)
// 		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
// 	}
// 	defer conn.Close()

// 	fmt.Println("Finished")
// }

// func rot13handle(conn net.Conn) {
// 	scanner := bufio.NewScanner(conn)
// 	for scanner.Scan() {
// 		ln := strings.ToLower(scanner.Text())
// 		bs := []byte(ln)
// 		r := rot13(bs)

// 		fmt.Fprintf(conn, "%s - %s\n\n", ln, r)
// 	}
// }

// func rot13(bs []byte) []byte {
// 	var r13 = make([]byte, len(bs))
// 	for i, v := range bs {
// 		if v <= 109 {
// 			r13[i] = v + 13
// 		} else {
// 			r13[i] = v - 13
// 		}
// 	}
// 	return r13
// }

func dbhandle(conn net.Conn) {
	defer conn.Close()

	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fs := strings.Fields(scanner.Text())
		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s\n", v)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "You need to provide a key value pair")
				continue
			}
			data[fs[1]] = fs[2]
		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintln(conn, "I don't know that command")
		}
	}
}
