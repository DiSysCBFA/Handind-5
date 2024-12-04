package client

import (
	"context"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"

	h5 "github.com/DiSysCBFA/Handind-5/Api"
	"github.com/manifoldco/promptui"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartClient(ports []string, id string) {
	conn, err := grpc.Dial(ports[0], grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := h5.NewAuctionserviceClient(conn)
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
	for i, port := range ports {
		log.Println("Dialing", port)
		conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("Failed to dial port %s: %v", port, err)
			continue
		}
		defer conn.Close()

		empty := &h5.Empty{}

		client := h5.NewAuctionserviceClient(conn)
		response, err := client.JoinAuction(context.Background(), empty)
		if err != nil {
			log.Printf("Failed to join auction: %v", err)
			continue
		}
		log.Println("Current status:", response.Status)

		responses[i] = response
	}

	if responses[0] != nil && responses[1] != nil && responses[2] != nil {
		if reflect.DeepEqual(responses[0], responses[1]) {
			log.Println(responses[0].Status)
		} else if reflect.DeepEqual(responses[0], responses[2]) {
			log.Println(responses[0].Status)
		} else if reflect.DeepEqual(responses[1], responses[2]) {
			log.Println(responses[1].Status)
		} else {
			log.Println("Servers don't agree. Øv Bøv")
		}
	} else {
		log.Println("Not all responses are valid")
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

	responses := make([]*h5.BidAck, len(ports))
	for index, port := range ports {
		conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("The server at port %s has crashed: %v", port, err)
			responses[index] = nil // Explicitly set response to nil
			continue
		}
		defer conn.Close()

		activeBid := &h5.Bid{
			Bid:       int64(intBid),
			Bidder:    id,
			Timestamp: time.Now().UnixNano(),
		}

		client = h5.NewAuctionserviceClient(conn)
		respons, err := client.TryBid(context.Background(), activeBid)
		if err != nil {
			log.Printf("Failed to submit bid to server at port %s: %v", port, err)
			responses[index] = nil
			continue
		}
		responses[index] = respons
	}

	// Ensure responses are not nil before comparison
	for i := 0; i < len(responses); i++ {
		if responses[i] == nil {
			log.Printf("No response from server at port %s", ports[i])
			continue
		}
		log.Printf("Response from server at port %s: %s", ports[i], responses[i].Ack)
	}
}
