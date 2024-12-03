package client

import (
	"context"
	"log"
	"os"
	"reflect"
	"strconv"
	"sync"
	"time"

	h5 "github.com/DiSysCBFA/Handind-5/Api"
	"github.com/manifoldco/promptui"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartClient(wg *sync.WaitGroup, ports []string, id string) {
	defer wg.Done()
	client := h5.NewAuctionserviceClient(nil)
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
			getBid(client, id, ports)
		} else if selectedAction == "Result" {
			retreiveResult(ports)
		} else if selectedAction == "Exit" {
			log.Println("Exiting...")
			os.Exit(0)
		}
	}
}

func retreiveResult(ports []string) {
	responses := [3]*h5.AuctionResult{}
	for _, port := range ports {
		log.Println("Dialing", port)
		conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Failed to dial: %v. Attempting next port", err)
			continue
		}

		empty := &h5.Empty{}

		client := h5.NewAuctionserviceClient(conn)
		response, err := client.JoinAuction(context.Background(), empty)
		if err != nil {
			log.Fatalf("Failed to join auction: %v", err)
			continue
		}
		log.Println("Current status:", response.Status)

		if reflect.DeepEqual(responses[0], responses[1]) {
			log.Println(responses[0].Status)
		} else if reflect.DeepEqual(responses[0], responses[2]) {
			log.Println(responses[0].Status)
		} else if reflect.DeepEqual(responses[1], responses[2]) {
			log.Println(responses[1].Status)
		} else {
			log.Println("Servers don't agree. Øv Bøv")
		}
	}
}

func getBid(client h5.AuctionserviceClient, id string, ports []string) {
	enterBid := promptui.Prompt{
		Label: "Enter bid",
	}
	bid, err := enterBid.Run()
	if err != nil {
		log.Fatalf("Failed to enter bid: %v", err)
	}
	intBid, err := strconv.Atoi(bid)
	if err != nil {
		log.Fatalf("Failed to convert bid to int: %v", err)
	}

	responses := [3]*h5.BidAck{}

	for index, port := range ports {
		conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Println("The server at port " + port + " has crashed")
			continue
		}
		activeBid := &h5.Bid{
			Bid:       int64(intBid),
			Bidder:    id,
			Timestamp: time.Now().Unix(),
		}

		client = h5.NewAuctionserviceClient(conn)
		respons, err := client.TryBid(context.Background(), activeBid)
		if err != nil {
			log.Println("The server at port " + port + " has crashed")
			continue
		}
		responses[index] = respons
	}
	if reflect.DeepEqual(responses[0], responses[1]) {
		log.Println(responses[0].Ack)
	} else if reflect.DeepEqual(responses[0], responses[2]) {
		log.Println(responses[0].Ack)
	} else if reflect.DeepEqual(responses[1], responses[2]) {
		log.Println(responses[1].Ack)
	} else {
		log.Println("Servers cant reach consensus. More than one server dont work")
	}
}
