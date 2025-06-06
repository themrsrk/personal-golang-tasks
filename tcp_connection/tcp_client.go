package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	for {
		var message string
		fmt.Println("Enter your message: ")
		fmt.Scan(&message)

		//Send message to server
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		fmt.Println("Connected to server ...")
	}

}
