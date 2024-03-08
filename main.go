package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/fajarlubis/meatball/server"

	_ "net/http/pprof"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "pong")
		default:
			http.NotFound(w, r)
		}
	})

	s := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	go server.Start("2271")
	go server.Start("2273")
	go server.Start("2275")

	log.Println(runtime.NumGoroutine())

	log.Fatal(s.ListenAndServe())
}
