package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", plainText)
	http.ListenAndServe(":3000", nil)
	fmt.Print("listening on port 3000")

}
func writeHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("server", "server-1")
	w.WriteHeader(200)
}

func plainText(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
