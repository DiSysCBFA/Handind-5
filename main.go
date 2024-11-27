package main

import (
	"bufio"
	client "github.com/DiSysCBFA/Handind-5/Client"
	"log"
	"os"
	"strings"

	_ "github.com/DiSysCBFA/Handind-5/Server"
	
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
			bidder, err := selectBidderName.Run()
			if err != nil {
				log.Fatalf("Failed to run: %v", err)
			}
			log.Println("Bidder name:", bidder)
			b := client.NewBidder(bidder, "4000")
			go b.Join([]string{"localhost:4000", "localhost:4001", "localhost:4002"})
			go b.SendBids([]string{"localhost:4000", "localhost:4001", "localhost:4002"})

		} else if result == "Exit" {
			log.Println("Exiting...")
			os.Exit(0)
		}
	}
}
