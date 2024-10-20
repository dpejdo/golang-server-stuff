package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn net.Conn
	name string
}

var clients = make(map[net.Conn]string)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Chat server is running on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("New client connected:", conn.RemoteAddr())
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	fmt.Println("Sending name prompt to client:", conn.RemoteAddr())
	_, err := writer.WriteString("Enter your name: ")
	if err != nil {
		fmt.Println("Error writing to client:", err)
		return
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}

	fmt.Println("Waiting for client name:", conn.RemoteAddr())
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading client name:", err)
		return
	}
	name = strings.TrimSpace(name)
	fmt.Println("Client name received:", name)

	clients[conn] = name
	broadcast(fmt.Sprintf("%s has joined the chat", name))

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			delete(clients, conn)
			broadcast(fmt.Sprintf("%s has left the chat", name))
			fmt.Println("Client disconnected:", name)
			return
		}

		message = strings.TrimSpace(message)
		if message != "" {
			fmt.Printf("Message from %s: %s\n", name, message)
			broadcast(fmt.Sprintf("%s: %s", name, message))
		}
	}
}

func broadcast(message string) {
	fmt.Println("Broadcasting:", message)
	for conn := range clients {
		writer := bufio.NewWriter(conn)
		_, err := writer.WriteString(message + "\n")
		if err != nil {
			fmt.Println("Error broadcasting to client:", err)
			continue
		}
		err = writer.Flush()
		if err != nil {
			fmt.Println("Error flushing writer:", err)
			continue
		}
	}
}
