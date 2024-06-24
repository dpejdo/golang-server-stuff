package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	fmt.Println("listening on port")

	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v\n%v", w, r)
	w.Header().Set("Access-Control-Allow-Origin", "*") // not a real fix for cors
	// more of quick fix
	// cors should be set to specific origin where is your production running
	// but this is on localhost so it's does not matter but it's really important for production
	io.WriteString(w, "hello world!")

}
