package main

import (
	"NetKnock/Colors"
	"fmt"
	"net"
	"os"
	"time"
)

var color = ""

func currentTime() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("02/01/2006 15:04:05")
	return string(formattedTime)
}

func toggleColor() string {
	if color == "" {
		color = Colors.YELLOW
		return color
	}
	if color == Colors.YELLOW {
		color = Colors.CYAN
		return color
	}
	color = Colors.YELLOW
	return color
}

func main() {
	// Check if the address and port are provided as an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: netknock <address:port> (without http:// or https://")
		os.Exit(1)
	}

	// Get the address and port from the first argument
	address := os.Args[1]
	fmt.Printf("──────────────────────────────────────────────────────────────────────────────────────────\n")
	fmt.Printf("│" + Colors.YELLOW + "  NetKnock by AndersonPEM <https://github.com/andersonpem https://gitlab.com/andersonpem>" + Colors.RESET + "│\n")
	fmt.Printf("│" + Colors.GREEN + "One Knock to Rule Them All." + Colors.RESET + "                                                              │\n")
	fmt.Printf("│" + Colors.GREEN + "I'll be knocking your port until it responds. And I'm very insistent. ^^" + Colors.RESET + "                 │\n")
	fmt.Printf("│"+Colors.CYAN+"Host to be checked: %s"+Colors.RESET+"                                                       │\n", address)
	fmt.Printf("──────────────────────────────────────────────────────────────────────────────────────────\n")

	for {
		// Attempt to connect to the address and port
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// Log the error if the connection fails
			fmt.Printf(toggleColor()+"["+currentTime()+"]: I still cannot connect to %s. Error: %s.\nWill retry in a second...\n"+Colors.RESET, address, err)
		} else {
			// Log success and close the connection
			fmt.Printf(Colors.GREEN+"Successfully connected to %s\n! We're good to go!"+Colors.RESET, address)
			err := conn.Close()
			if err != nil {
				fmt.Printf(Colors.RED+"Some unexpected error happened.\n%s\nWill exit now."+Colors.RESET, err.Error())
				os.Exit(1)
			}
			// Exit with status code 0
			os.Exit(0)
		}

		// Wait for 1 second before attempting again
		time.Sleep(1 * time.Second)
	}
}
