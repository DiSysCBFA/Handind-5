package Server

import (
	"context"
	"fmt"
	"log"
	"time"

	api "github.com/DiSysCBFA/Handind-5/Api"
)

type Server struct {
	api.UnimplementedAuctionserviceServer
	Port                 string
	CurrentHighestBid    int64
	CurrentHighestBidder string
	AuctionStatus        string
	Bidders              map[string]int64
}

func NewServer(port string) *Server {
	return &Server{
		Port:                 port,
		CurrentHighestBid:    0,
		CurrentHighestBidder: "",
		AuctionStatus:        "Active",
		Bidders:              make(map[string]int64),
	}
}

func (s *Server) SendBid(ctx context.Context, bid *api.Bid) (*api.BidAck, error) {

	if s.AuctionStatus == "ENDED" {
		return &api.BidAck{Ack: "Auction has ended. No more bids allowed."}, nil // Checking if the auction is active
	}

	if bid.Bid <= s.CurrentHighestBid { // checking the bid amount
		return &api.BidAck{Ack: fmt.Sprintf("Bid too low. Current highest bid is %d.", s.CurrentHighestBid)}, nil
	}

	if _, exists := s.Bidders[bid.Bidder]; !exists { // Register the bidder if it's their first bid
		s.Bidders[bid.Bidder] = bid.Bid
	}
	s.CurrentHighestBid = bid.Bid
	s.CurrentHighestBidder = bid.Bidder
	log.Printf("New highest bid: %d by %s", bid.Bid, bid.Bidder) // Update the auction state

	return &api.BidAck{Ack: "Bid accepted"}, nil // Returns "Bid accepted" if success
}
func (s *Server) JoinAuction(empty *api.Empty, stream api.Auctionservice_JoinAuctionServer) error { // Streaming out the current auction state periodically
	for {
		if s.AuctionStatus == "ENDED" { // Check if the auction has ended
			return stream.Send(&api.Auction{
				Status: fmt.Sprintf("Auction ended. Winner: %s with bid %d", s.CurrentHighestBidder, s.CurrentHighestBid),
			})
		}
		err := stream.Send(&api.Auction{
			Status: fmt.Sprintf("Current highest bid: %d by %s", s.CurrentHighestBid, s.CurrentHighestBidder), // Streaming the current highest bid and bidder
		})
		if err != nil {
			log.Printf("Error sending auction update to client: %v", err) // Error if can't stream
			return err
		}
		time.Sleep(2 * time.Second)
	}

}
