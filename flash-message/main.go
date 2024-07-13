package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/get", get)
	http.HandleFunc("/set", set)
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}

func get(w http.ResponseWriter, r *http.Request) {

	value, err := getFlash(w, r, "msg")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if value == nil {
		fmt.Fprintf(w, "Empty message")
		return
	}

	fmt.Fprintf(w, "%s", value)
}

func set(w http.ResponseWriter, r *http.Request) {
	setFlash(w, "msg", "this is a message")

}
