package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

func main() {

	http.HandleFunc("/get", get)
	http.HandleFunc("/set", set)
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}

var store = sessions.NewCookieStore([]byte("a secret string"))

func set(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "flash-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.AddFlash("This is a flashed message!", "message")
	session.Save(r, w)
}

func get(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "flash-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fm := session.Flashes("message")
	if fm == nil {
		fmt.Fprint(w, "No flash messages")
		return
	}

	session.Save(r, w)
	fmt.Fprintf(w, "%v", fm[0])
}
