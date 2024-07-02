package main

import (
	"log"
	"net/http"
)

func middlewearOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing middlewear one")
		next.ServeHTTP(w, r)
		log.Print("Again executing middlewear one")
	})
}

func middlewearTwo(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing middlewear two")
		if r.URL.Path == "/foo" {
			return
		}

		next.ServeHTTP(w, r)
		log.Print("Again executing middleear two")
	})

}

func finalHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("executing final handler")
	w.Write([]byte("ok"))

}
func main() {
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(finalHandler)
	mux.Handle("/", middlewearOne((middlewearTwo(finalHandler))))
	http.ListenAndServe(":3000", mux)

}
