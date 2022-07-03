package main

import (
	"fmt"
	"os"
)

func main() {
	intro()
	printMenu()

	command := readCommand()

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Displaying Logs")
	case 0:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}
}

func intro() {
	name := "Maverick" // equivalent to 'var name string = "Maverick'"
	version := 1.0     // Equivalent to 'var version float64 = 1.0'

	fmt.Println("Hello, sr", name)
	fmt.Println("This software is running version", version)
}

func printMenu() {
	fmt.Println("\n1. Start monitoring")
	fmt.Println("2. Show logs")
	fmt.Println("0. Exit")
	fmt.Print("> ")
}

func readCommand() int {
	var command int

	fmt.Scan(&command) // Scan can infer variable type. Ignore value if are not expected

	fmt.Println("\nThe chosen command was:", command)

	return command
}
