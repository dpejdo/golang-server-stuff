package main

import (
	"log"
	"net/http"

	"github.com/goji/httpauth"
)

func final(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	authHandler := httpauth.SimpleBasicAuth("alice", "password")
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", authHandler(finalHandler))

	log.Print("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
