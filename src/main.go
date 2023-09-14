package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// Check if the address and port are provided as an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: netknock <address:port> (without http:// or https://")
		os.Exit(1)
	}

	// Get the address and port from the first argument
	address := os.Args[1]
	fmt.Printf("Waiting for %s to be available...", address)

	for {
		// Attempt to connect to the address and port
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// Log the error if the connection fails
			fmt.Printf("Failed to connect to %s. Error: %s\n. Will retry in a second...\n", address, err)
		} else {
			// Log success and close the connection
			fmt.Printf("Successfully connected to %s\n! We're good to go!", address)
			err := conn.Close()
			if err != nil {
				return
			}
			// Exit with status code 0
			os.Exit(0)
		}

		// Wait for 1 second before attempting again
		time.Sleep(1 * time.Second)
	}
}
