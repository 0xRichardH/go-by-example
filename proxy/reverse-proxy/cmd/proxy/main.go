package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	originalServerURL, err := url.Parse("http://localhost:3000/hey")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/hey", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("reverse proxy request: %s", r.URL.String())

		r.Host = originalServerURL.Host
		r.URL.Host = originalServerURL.Host
		r.URL.Scheme = originalServerURL.Scheme
		r.RequestURI = ""
		r.Header.Add("X-Forwarded-For", r.RemoteAddr)

		resp, err := http.DefaultClient.Do(r)
		if err != nil {
			log.Printf("failed to send request: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		w.WriteHeader(http.StatusOK)
		if _, err := io.Copy(w, resp.Body); err != nil {
			log.Printf("failed to copy response: %v", err)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
