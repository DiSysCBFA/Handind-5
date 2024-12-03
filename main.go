package main

import (
	"log"
	"os"
	"sync"

	client "github.com/DiSysCBFA/Handind-5/Client"
	"github.com/manifoldco/promptui"
)

func main() {
	var wg sync.WaitGroup
	ports := []string{"localhost:4000", "localhost:4001", "localhost:4002"}
	//maxRetries := len(ports) - 1

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
			/* auctionServer := server.NewServer("")
			auctionServer.Start(ports[0], maxRetries) */
		} else if result == "Start Client" {

			enterID := promptui.Prompt{}
			enteredID, err := enterID.Run()
			if err != nil {
				log.Fatalf("Failed to enter ID: %v", err)
			}
			go client.StartClient(&wg, ports, enterID)
		} else if result == "Exit" {
			log.Println("Exiting...")
			os.Exit(0)
		}
		wg.Wait()
	}

}
