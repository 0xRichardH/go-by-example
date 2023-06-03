package main

import (
	"io"
	"log"
	"net/http"
)

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("proxy: %s", r.URL)
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		log.Printf("error: failed to make request: %v", err)
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	if _, err := io.Copy(w, resp.Body); err != nil {
		log.Printf("failed to copy response: %v", err)
	}
}

func copyHeader(dst, src http.Header) {
	for k, vArr := range src {
		for _, v := range vArr {
			dst.Add(k, v)
		}
	}
}
