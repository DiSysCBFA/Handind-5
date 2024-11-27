package client

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"time"

	h5 "github.com/DiSysCBFA/Handind-5/Api"
	"google.golang.org/grpc"

	auction "github.com/DiSysCBFA/Handind-5/Api"
)

// Bidder represents an auction client with a gRPC connection
type Bidder struct {
	auction.AuctionserviceClient
	conn      *grpc.ClientConn
	port      string
	bidder    string
	timestamp int
}

// NewBidder creates a new client instance and initializes the gRPC connection
func NewBidder(name, port string) *Bidder {
	address := "localhost:" + port
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	client := auction.NewAuctionserviceClient(conn)
	return &Bidder{
		AuctionserviceClient: client,
		conn:                 conn,
		port:                 port,
		bidder:               name,
		timestamp:            0,
	}
}

func (b *Bidder) Join(ports []string) {
	for i, port := range ports {
		conn, err := grpc.Dial(port, grpc.WithInsecure())
		if err == nil {
			defer conn.Close()
			client := h5.NewAuctionserviceClient(conn)
			stream, err := client.JoinAuction(context.Background(), &auction.Empty{})
			if err == nil {
				go func() {
					for {
						in, err := stream.Recv()
						if err != nil {
							log.Printf("Failed to receive auction update: %v", err)
							break
						}
						log.Printf("Auction status: %s", in.Status)
					}
				}()
				return
			}
		}
		if i == len(ports)-1 {
			log.Println("All servers are down")
			return
		}
	}
}

// SendBids prompts the user to send bids
func (b *Bidder) SendBids(ports []string) {
	reader := bufio.NewReader(os.Stdin)
	var bidamount int64
	fmt.Print("Enter your bid: ")
	_, err := fmt.Fscanf(reader, "%d\n", &bidamount)
	if err != nil {
		log.Fatal(err)
	}
	newbid := &h5.Bid{
		Bidder:    b.bidder,
		Bid:       bidamount,
		Timestamp: time.Now().Unix(),
	}

	for _, port := range ports {
		go func(port string) {
			conn, err := grpc.Dial(port, grpc.WithInsecure())
			if err != nil {
				log.Printf("Failed to connect to %s: %v", port, err)
				return
			}
			defer func(conn *grpc.ClientConn) {
				err := conn.Close()
				if err != nil {

				}
			}(conn)

			client := h5.NewAuctionserviceClient(conn)
			_, err = client.SendBid(context.Background(), newbid)
			if err != nil {
				log.Printf("Error sending bid to auction")
			} else {
				log.Printf("bid sent to auction")

			}
		}(port)
	}
}

// Close closes the gRPC connection
func (b *Bidder) Close() {
	if b.conn != nil {
		b.conn.Close()
	}
}
