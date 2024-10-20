package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Attempting to connect to the server...")
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to server successfully.")

	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		if err == io.EOF {
			fmt.Println("Server closed the connection")
		} else if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			fmt.Println("Timeout waiting for data from server")
		} else {
			fmt.Printf("Error reading from server: %v\n", err)
		}
	} else {
		fmt.Printf("Received %d bytes from server: %s\n", n, string(buffer[:n]))
	}

	conn.SetReadDeadline(time.Time{})

	name, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading name:", err)
		return
	}
	name = strings.TrimSpace(name)

	fmt.Println("Sending name to server:", name)
	_, err = fmt.Fprintf(conn, "%s\n", name)
	if err != nil {
		fmt.Println("Error sending name to server:", err)
		return
	}

	go receiveMessages(conn)
	sendMessages(conn)
}

func receiveMessages(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println("Received:", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from server:", err)
	}
}

func sendMessages(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println("Sending message:", message)
		_, err := fmt.Fprintf(conn, "%s\n", message)
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from stdin:", err)
	}
}
