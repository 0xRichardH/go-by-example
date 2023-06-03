package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hey", func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.Header.Get("X-Forwarded-For")
		if len(clientIP) > 0 {
			log.Printf("new request from %s", clientIP)
		}
		fmt.Fprintln(w, "Hey, world")
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
