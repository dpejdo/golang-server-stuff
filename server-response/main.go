package main

import (
	"net/http"
	"path"
	"text/template"
)

func main() {
	http.HandleFunc("/", serverHtmlTemplate)
	http.ListenAndServe(":3000", nil)

}

type Person struct {
	Name    string
	Hobbies []string
}

func serverHtmlTemplate(w http.ResponseWriter, r *http.Request) {
	person := Person{Name: "John", Hobbies: []string{"coding", "reading"}}

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
