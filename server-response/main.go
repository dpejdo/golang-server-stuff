package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", writeJson)
	http.ListenAndServe(":3000", nil)
	fmt.Print("listening on port 3000")

}

type Profile struct {
	Name    string
	Hobbies []string
}

func writeJson(w http.ResponseWriter, r *http.Request) {
	profile := Profile{Name: "john", Hobbies: []string{"reading", "coding"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(js)
}
