package main

import (
	"bufio"
	client "github.com/DiSysCBFA/Handind-5/Client"
	"log"
	"net"
	"os"
	"strings"
	"github.com/manifoldco/promptui"

	api "github.com/DiSysCBFA/Handind-5/Api"
	server "github.com/DiSysCBFA/Handind-5/Server"
	"google.golang.org/grpc"

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

	for i := 0; i < 3; i++ {
		port, err := r.ReadString('\n') //! Make sure last line has a new line
		log.Println(port)

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
			bidder, err := selectBidderName.Run()
			if err != nil {
				log.Fatalf("Failed to run: %v", err)
			}
			log.Println("Bidder name:", bidder)
			b := client.NewBidder(bidder, "4000")
			go b.Join([]string{"localhost:4000", "localhost:4001", "localhost:4002"})
			go b.SendBids([]string{"localhost:4000", "localhost:4001", "localhost:4002"})

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
