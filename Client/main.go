package client

import (
	"log"
	"os"
	"sync"

	"github.com/manifoldco/promptui"
)

func StartClient(wg *sync.WaitGroup, ports []int) {
	defer wg.Done()
	for {
		selectAction := promptui.Select{
			Label: "Select an option",
			Items: []string{"Bid", "Result", "Exit"},
		}
		_, selectedAction, err := selectAction.Run()
		if err != nil {
			log.Fatalf("Failed to run selection: %v", err)
		}

		if selectedAction == "Bid" {
		} else if selectedAction == "Result" {
		} else if selectedAction == "Exit" {
			log.Println("Exiting...")
			os.Exit(0)
		}
	}
}
