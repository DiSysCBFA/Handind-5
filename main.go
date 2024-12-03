package main

import (
	"log"

	server "github.com/DiSysCBFA/Handind-5/Server"
	"github.com/manifoldco/promptui"
)

func main() {
	ports := []int{4000, 4001, 4002, 4003, 4004}
	maxRetries := len(ports) - 1

	for {
		selection := promptui.Select{
			Label: "Select an option",
			Items: []string{"Start Server", "Start Client", "Exit"},
		}

		_, result, err := selection.Run()
		if err != nil {
			log.Fatalf("Failed to run selection: %v", err)
		}

		if result == "Start Server" {
			log.Println("Starting server...")
			auctionServer := server.NewServer("")
			auctionServer.Start(ports[0], maxRetries)
		} else if result == "Start Client" {

		} else if result == "Exit" {
			log.Println("Exiting...")
			break
		}
	}
}
