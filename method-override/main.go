package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

const form = `
<!DOCTYPE HTML>
<html>
    <body>
        <form method="POST" action="/">
            <input type="hidden" name="_method" value="PUT">
            <label>Example field</label>
            <input type="text" name="example">
            <button type="submit">Submit</button>
        </form>
    </body>
</html>
`

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", formHandler)

	err := http.ListenAndServe(":3000", methodOverride(mux))
	log.Print(err)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t, err := template.New("form").Parse(form)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}

		t.Execute(w, nil)

	case http.MethodPut:
		io.WriteString(w, "this is put method")
	default:
		http.Error(w, http.StatusText(405), 405)
	}
}
