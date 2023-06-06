package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	startServer()
}

func startServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	// for {
	// 	conn, err := ln.Accept()
	// 	if err != nil {
	// 		log.Printf("error accepting: %s", err)
	// 	}
	// 	go handlerServerConnection(conn)
	// }

	conn, err := ln.Accept()
	if err != nil {
		log.Printf("error accepting: %s", err)
		return
	}
	handlerServerConnection(conn)
}

func handlerServerConnection(conn net.Conn) {
	defer conn.Close()

	r, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Printf("failed to read: %s", err)
		return
	}
	log.Printf("read: %s", r)

	respBody := fmt.Sprintf("time: %s \r\n", time.Now().Format(time.RFC3339))
	if _, err := fmt.Fprint(conn, respBody); err != nil {
		log.Printf("failed to write: %s", err)
	}
}
