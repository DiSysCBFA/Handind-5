package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

func main() {
	file, err := os.Open("ports.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	r := bufio.NewReader(file)

	nop, err := r.ReadString('\n')
	var numberOfPeers, _ = strconv.Atoi(strings.TrimSpace(nop))
	log.Println("Number of peers is: ", numberOfPeers)
	if err != nil {
		log.Fatal(err)
		return
	}
	for i := 0; i < numberOfPeers; i++ {
		port, err := r.ReadString('\n') //! Make sure last line has a new line
		log.Println(port)
		if err != nil {
			break
		}

		port = strings.TrimSpace(port)
		log.Println(port) // Displaying each port read

		selection := promptui.Select{
			Label: "Select an option",
			Items: []string{"Start Server", "Start new Client", "Exit"},
		}

		_, result, err := selection.Run()
		if err != nil {
			log.Fatalf("Failed to run: %v", err)
		}

		if result == "Start Server" {
			// Start the server
		} else if result == "Start new Client" {
			// Start a new client
		} else if result == "Exit" {
			log.Println("Exiting...")
			os.Exit(0)
		}
	}
}
