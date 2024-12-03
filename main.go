package main

import (
	"log"
	"os"

	client "github.com/DiSysCBFA/Handind-5/Client"
	server "github.com/DiSysCBFA/Handind-5/Server"
	"github.com/manifoldco/promptui"
)

func main() {
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
			server.StartServer()

		} else if result == "Start Client" {

			enterID := promptui.Prompt{}
			enteredID, err := enterID.Run()
			if err != nil {
				log.Fatalf("Failed to enter ID: %v", err)
			}
			client.StartClient(ports, enteredID)
		} else if result == "Exit" {
			log.Println("Exiting...")
			os.Exit(0)
		}
	}

}
