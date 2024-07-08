package main

import (
	"log"
	"net/http"

	"github.com/bmizerany/pat"
)

func main() {
	mux := pat.New()

	mux.Get("/:locale", http.HandlerFunc(handleHome))

	log.Print("Listening on port :3000")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
