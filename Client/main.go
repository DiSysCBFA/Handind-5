package Client

import (
	"context"
	"log"
	"time"

	pb "github.com/DiSysCBFA/Handind-5/Api" // Import generated proto package
	"google.golang.org/grpc"
)

func main() {
	// Connect to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAuctionServiceClient(conn)

	// Make a bid
	bidResponse, err := client.Bid(context.Background(), &pb.BidRequest{
		Bidder: "Alice",
		Amount: 100,
	})
	if err != nil {
		log.Fatalf("Error calling Bid: %v", err)
	}
	log.Printf("Bid Response: %v", bidResponse.Status)

	// Query result
	time.Sleep(2 * time.Second) // Wait for the auction duration to pass
	resultResponse, err := client.Result(context.Background(), &pb.ResultRequest{})
	if err != nil {
		log.Fatalf("Error calling Result: %v", err)
	}
	log.Printf("Auction Result: Bidder=%s, Amount=%d, Closed=%v",
		resultResponse.HighestBidder, resultResponse.HighestBid, resultResponse.IsAuctionClosed)
}
