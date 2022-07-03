package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const CHECK_NUMBER = 3
const CHECK_DELAY = 5

func main() {

	intro()

	for {
		printMenu()

		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
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

func startMonitoring() {
	fmt.Println("Monitoring...")

	sites := readSitesFromFile()

	for i := 0; i < CHECK_NUMBER; i++ {
		fmt.Println("> Run", i+1)

		for _, site := range sites {
			healthCheck(site)
		}

		// Skip sleep on last run
		if i < CHECK_NUMBER-1 {
			fmt.Println("\nDelaying next check for", CHECK_DELAY, "seconds\n")
			time.Sleep(CHECK_DELAY * time.Second)
		}
	}

}

func healthCheck(site string) {
	res, err := http.Get(site)

	if err != nil {
		fmt.Println("Site health check failed", err)
	}

	if res.StatusCode == 200 {
		fmt.Println(site, "is up")
	} else {
		fmt.Println(site, "is down. Status Code:", res.StatusCode)
	}
}

func readSitesFromFile() []string {
	sites := []string{}

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Failure when opening file: ", err)
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}
