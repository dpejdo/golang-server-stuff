package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("error happends while starting server %+v ", err)
	}

	log.Println("started on port 8080")
	conn, err := listener.Accept()
	if err != nil {
		log.Printf("error while accepting connection %+v ", err)
		return
	}

	var data []byte
	_, err = conn.Read(data)

	if err != nil {
		log.Printf("Error while parsing data %+v", err)
		return
	}

	log.Println(string(data))

	conn.Write([]byte("hello"))
	conn.Close()

}
