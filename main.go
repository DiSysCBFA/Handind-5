package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"

	"github.com/manifoldco/promptui"

	api "github.com/DiSysCBFA/Handind-5/Api"
	client "github.com/DiSysCBFA/Handind-5/Client"
	server "github.com/DiSysCBFA/Handind-5/Server"
	"google.golang.org/grpc"
)

func main() {
	// Open the ports file
	file, err := os.Open("ports.txt")
	if err != nil {
		log.Fatalf("Failed to open ports file: %v", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// Loop for handling up to 3 iterations (3 servers/clients)
	for i := 0; i < 3; i++ {
		port, err := reader.ReadString('\n') // Ensure ports.txt has newline-delimited ports
		if err != nil {
			log.Printf("Error reading port: %v", err)
			break
		}

		port = strings.TrimSpace(port) // Remove unnecessary whitespace/newline
		log.Printf("Reading port: %s", port)

		// Prompt user to select an option
		selection := promptui.Select{
			Label: "Select an option",
			Items: []string{"Start Server", "Start New Client", "Exit"},
		}

		_, result, err := selection.Run()
		if err != nil {
			log.Fatalf("Failed to run selection: %v", err)
		}

		if result == "Start Server" {
			// Start the server
			log.Printf("Starting server on port %s...", port)
			go startServer(port)

		} else if result == "Start New Client" {
			// Prompt for bidder name
			selectBidderName := promptui.Prompt{
				Label: "Enter desired bidder name",
			}
			bidderName, err := selectBidderName.Run()
			if err != nil {
				log.Fatalf("Failed to get bidder name: %v", err)
			}

			log.Printf("Bidder name: %s", bidderName)

			// Create and start a new client
			bidder := client.NewBidder(bidderName, port)
			log.Printf("Starting bidder %s on port %s...", bidderName, port)

			go bidder.Join([]string{"localhost:4000", "localhost:4001", "localhost:4002"})
			go bidder.SendBids([]string{"localhost:4000", "localhost:4001", "localhost:4002"})

		} else if result == "Exit" {
			// Exit the application
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
	auctionServer := server.NewServer(port)                     // Initialize a new server instance
	api.RegisterAuctionServiceServer(grpcServer, auctionServer) // Register the gRPC service

	log.Printf("Server started on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
