package main

import (
	"log"
	"net/http"

	"github.com/seansu4you87/gocached/server"
)

const addr = "localhost:1988"

func main() {
	mux := http.NewServeMux()
	indexHandler := &server.IndexHandler{}
	mux.Handle("/", indexHandler)

	log.Printf("Now listening on %s...\n", addr)

	server := http.Server{Handler: mux, Addr: addr}
	log.Fatal(server.ListenAndServe())
}
