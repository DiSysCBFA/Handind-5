package main

import (
	"bufio"
	"log"
	"os"
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

	for i := 0; i < 3; i++ {
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
			selectBidderName := promptui.Prompt{
				Label: "Enter desired name",
			}
			Bidder, err := selectBidderName.Run()
			if err != nil {
				log.Fatalf("Failed to run: %v", err)
			}
			log.Println("Bidder name:", Bidder)

		} else if result == "Exit" {
			log.Println("Exiting...")
			os.Exit(0)
		}
	}
}
