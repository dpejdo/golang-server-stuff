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
	io.WriteString(w, "hello world!")

}
