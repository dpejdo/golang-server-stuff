package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandler(t string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(t)
		w.Write([]byte("Time is: " + tm))
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/time", timeHandler(time.RFC1123))
	log.Print("Listening ....")
	http.ListenAndServe(":3000", mux)
}
