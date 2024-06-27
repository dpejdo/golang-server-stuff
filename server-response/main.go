package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", writeXml)
	http.ListenAndServe(":3000", nil)
	fmt.Print("listening on port 3000")

}

type Profile struct {
	Name    string
	Hobbies []string
}

func writeXml(w http.ResponseWriter, r *http.Request) {
	profile := Profile{Name: "john", Hobbies: []string{"reading", "coding"}}

	js, err := xml.MarshalIndent(profile, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/xml")
	w.Write(js)
}
