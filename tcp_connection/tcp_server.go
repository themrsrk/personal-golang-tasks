package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen for incoming connections on port 8080
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}

	fmt.Println("Server listening on port 8000 ...")

	// Accept incoming connections and handle them
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	// Close the connection when we're done
	defer conn.Close()

	for {
		// Read incoming data
		buf := make([]byte, 2048)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Print the incoming data
		fmt.Printf("Received: %s \n", buf)
	}
}
