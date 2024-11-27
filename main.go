package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"

	api "github.com/DiSysCBFA/Handind-5/Api"
	server "github.com/DiSysCBFA/Handind-5/Server"
	"google.golang.org/grpc"
)

func main() {
	file, err := os.Open("ports.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	r := bufio.NewReader(file)

	nop, err := r.ReadString('\n')
	if err != nil {
		log.Fatal("Failed to read number of peers:", err)
		return
	}

	numberOfPeers, _ := strconv.Atoi(strings.TrimSpace(nop))
	log.Printf("Number of peers: %d", numberOfPeers)

	for i := 0; i < numberOfPeers; i++ {
		port, err := r.ReadString('\n')
		if err != nil {
			break
		}

		port = strings.TrimSpace(port)
		log.Printf("Reading port: %s", port)

		selection := promptui.Select{
			Label: "Select an option",
			Items: []string{"Start Server", "Start new Client", "Exit"},
		}

		_, result, err := selection.Run()
		if err != nil {
			log.Fatalf("Failed to run selection: %v", err)
		}

		if result == "Start Server" {
			log.Printf("Starting server on port %s...", port)
			startServer(port)
		} else if result == "Start new Client" {
			selectBidderName := promptui.Prompt{
				Label: "Enter desired name",
			}
			Bidder, err := selectBidderName.Run()
			if err != nil {
				log.Fatalf("Failed to run: %v", err)
			}
			log.Println("Bidder name:", Bidder)

			// Start a new client
		} else if result == "Exit" {
			log.Println("Exiting...")
			os.Exit(0)
		}
	}
}

func startServer(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}
	grpcServer := grpc.NewServer()
	auctionServer := server.NewServer(port)
	api.RegisterAuctionserviceServer(grpcServer, auctionServer)

	log.Printf("Server started on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
