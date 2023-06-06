package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"syscall"
	"time"
)

func main() {
	runClient()
}

func runClient() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for i := 0; i < 3; i++ {
		_, err := fmt.Fprint(conn, "GET / HTTP/1.0\r\n\r\n")
		if err != nil {
			if errors.Is(err, syscall.EPIPE) {
				log.Printf("broken pipe")
				return
			}
			log.Fatal(err)
		}

		// body, err := bufio.NewReader(conn).ReadString('\n')
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Printf("client: %s", body)

		time.Sleep(time.Second)
	}
}
