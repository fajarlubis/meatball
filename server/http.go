package server

import (
	"fmt"
	"log"
	"net/http"
)

func Start(port string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "pong from %s", port)
		default:
			http.NotFound(w, r)
		}
	})

	s := &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%s", port),
		Handler: mux,
	}

	log.Fatal(s.ListenAndServe())
}
