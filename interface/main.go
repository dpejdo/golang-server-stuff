package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
)

type Customer struct {
	FirstName string
	LastName  string
}

func (c *Customer) toJson(w io.Writer) error {
	js, err := json.Marshal(c)
	if err != nil {
		return err
	}

	_, err = w.Write(js)

	return err
}

func main() {
	customer := &Customer{"hello", "world"}
	var buf bytes.Buffer

	err := customer.toJson(&buf)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(buf.String())

	// write to file

	file, err := os.Create("./customer")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = customer.toJson(file)
	if err != nil {
		log.Fatal(err)
	}

}
