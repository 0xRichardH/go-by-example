package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gliderlabs/ssh"
	"github.com/go-chi/chi/v5"
	"github.com/oklog/ulid/v2"
)

var tunnelSessions sync.Map

func main() {
	go StartSSHServer()
	StartHttpServer()
}

type Tunnel struct {
	w      io.Writer
	doneCh chan struct{}
}

func NewTunnel(w io.Writer) *Tunnel {
	return &Tunnel{w: w, doneCh: make(chan struct{})}
}

func (s *Tunnel) GetWriter() io.Writer {
	return s.w
}

func (s *Tunnel) GetDoneCh() <-chan struct{} {
	return s.doneCh
}

func (s *Tunnel) Done() {
	close(s.doneCh)
}

func StartSSHServer() {
	ssh.Handle(func(s ssh.Session) {
		id := ulid.Make().String()
		fmt.Fprintf(s, "url: http://localhost:3000/%s\n", id)

		timeoutTimer := time.NewTimer(time.Minute)

		tunnelCh := make(chan *Tunnel)
		tunnelSessions.Store(id, tunnelCh)
		defer tunnelSessions.Delete(id)

		select {
		case <-timeoutTimer.C:
			fmt.Fprintln(s, "Timeout")
		case tunnel := <-tunnelCh:
			_, err := io.Copy(tunnel.w, s)
			if err != nil {
				log.Fatal(err)
			}
			tunnel.Done()
			fmt.Fprint(s, "We are done.\n")
		}
	})
	log.Fatal(ssh.ListenAndServe(":2222", nil))
}

func StartHttpServer() {
	r := chi.NewRouter()
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		sessionAny, ok := tunnelSessions.Load(id)
		if !ok {
			log.Printf("session %s not found", id)
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Session %s not found", id)
			return
		}
		tunnelCh, ok := sessionAny.(chan *Tunnel)
		if !ok {
			log.Print("Failed to assert session channel")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Internal server error")
			return
		}
		tunnel := NewTunnel(w)
		tunnelCh <- tunnel
		<-tunnel.GetDoneCh()
	})

	log.Fatal(http.ListenAndServe(":3000", r))
}
