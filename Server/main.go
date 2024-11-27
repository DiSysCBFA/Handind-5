package Server

import (
	"fmt"
	"log"

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
		Port: port,
	}
}

func (s *Server) SendBid(bid *api.Bid) (*api.BidAck, error) {

	// Checking if the auction is ongoing
	if s.AuctionStatus == "ENDED" {
		return &api.BidAck{Ack: "Auction has ended. No more bids allowed."}, nil
	}

	// checking the bid amount
	if bid.Bid <= s.CurrentHighestBid {
		return &api.BidAck{Ack: fmt.Sprintf("Bid too low. Current highest bid is %d.", s.CurrentHighestBid)}, nil
	}

	// Register the bidder if it's their first bid
	if _, exists := s.Bidders[bid.Bidder]; !exists {
		s.Bidders[bid.Bidder] = bid.Bid
	}

	// Update the auction state
	s.CurrentHighestBid = bid.Bid
	s.CurrentHighestBidder = bid.Bidder
	log.Printf("New highest bid: %d by %s", bid.Bid, bid.Bidder)

	// Return if success
	return &api.BidAck{Ack: "Bid accepted"}, nil
}

func (s *Server) JoinAuction(stream api.Auctionservice_JoinAuctionServer) error {
	return nil

}
