package client

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"

	auction "github.com/DiSysCBFA/Handind-5/Api"
)

// Client represents an auction client with a gRPC connection
type Client struct {
	auction.AuctionserviceClient
	conn      *grpc.ClientConn
	port      string
	name      string
	timestamp int
}

// NewClient creates a new client instance and initializes the gRPC connection
func NewClient(name, port string) *Client {
	address := "localhost:" + port
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	client := auction.NewAuctionserviceClient(conn)
	return &Client{
		AuctionserviceClient: client,
		conn:                 conn,
		port:                 port,
		name:                 name,
		timestamp:            0,
	}
}

// Join connects to the server and starts listening for auction updates on the JoinAuction stream
func (c *Client) Join() {
	// Start the JoinAuction stream to listen for incoming auction updates
	stream, err := c.AuctionserviceClient.JoinAuction(context.Background(), &auction.Empty{})
	if err != nil {
		log.Fatalf("Failed to join auction: %v", err)
	}

	// Listen for auction updates in a separate goroutine
	go func() {
		for {
			in, err := stream.Recv()
			if err != nil {
				log.Fatalf("Failed to receive auction update: %v", err)
			}
			log.Printf("Auction status: %s", in.Status)
		}
	}()

	// Start sending bids to the server
	c.SendBids()
}

// SendBids prompts the user to send bids
func (c *Client) SendBids() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter bid amount: ")
		bidAmount, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read bid amount: %v", err)
		}

		// Trim newline characters from the bid amount
		bidAmount = strings.TrimSpace(bidAmount)
		bid, err := strconv.ParseInt(bidAmount, 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse bid amount: %v", err)
		}
		// Send the bid to the server using the SendBid method
		_, err = c.AuctionserviceClient.SendBid(context.Background(), &auction.Bid{
			Bidder:    c.name,
			Bid:       bid,
			Timestamp: time.Now().Unix(),
		})
		if err != nil {
			log.Fatalf("Failed to send bid: %v", err)
		}
	}
}

// Close closes the gRPC connection
func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}
