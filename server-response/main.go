package main

import (
	"fmt"
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", serveFile)
	http.ListenAndServe(":3000", nil)
	fmt.Print("listening on port 3000")

}

func serveFile(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("images", "foo.png")
	http.ServeFile(w, r, fp)
}
