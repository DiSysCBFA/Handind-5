package Server

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	api "github.com/DiSysCBFA/Handind-5/Api"
	"google.golang.org/grpc"
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

func (s *Server) Start(initialPort int, maxRetries int) {
	port := initialPort

	var lis net.Listener
	var err error

	for i := 0; i <= maxRetries; i++ {
		address := fmt.Sprintf(":%d", port)
		lis, err = net.Listen("tcp", address)
		if err == nil {
			log.Printf("Server started on port %d", port)
			break
		}
		log.Printf("Port %d is unavailable, trying next port...", port)
		port++
	}

	if err != nil {
		log.Fatalf("Failed to start server after %d retries: %v", maxRetries, err)
		return
	}

	grpcServer := grpc.NewServer()
	api.RegisterAuctionserviceServer(grpcServer, s)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()
	log.Printf("Auction server listening on port %d", port)
	select {} // Keep the server running
}

func (s *Server) SendBid(ctx context.Context, bid *api.Bid) (*api.BidAck, error) {
	if s.AuctionStatus == "ENDED" {
		return &api.BidAck{Ack: "Auction has ended. No more bids allowed."}, nil
	}

	if bid.Bid <= s.CurrentHighestBid {
		return &api.BidAck{Ack: fmt.Sprintf("Bid too low. Current highest bid is %d.", s.CurrentHighestBid)}, nil
	}

	if _, exists := s.Bidders[bid.Bidder]; !exists {
		s.Bidders[bid.Bidder] = bid.Bid
	}
	s.CurrentHighestBid = bid.Bid
	s.CurrentHighestBidder = bid.Bidder
	log.Printf("New highest bid: %d by %s", bid.Bid, bid.Bidder)

	return &api.BidAck{Ack: "Bid accepted"}, nil
}

func (s *Server) JoinAuction(empty *api.Empty, stream api.Auctionservice_JoinAuctionServer) error {
	for {
		if s.AuctionStatus == "ENDED" {
			return stream.Send(&api.Auction{
				Status: fmt.Sprintf("Auction ended. Winner: %s with bid %d", s.CurrentHighestBidder, s.CurrentHighestBid),
			})
		}
		err := stream.Send(&api.Auction{
			Status: fmt.Sprintf("Current highest bid: %d by %s", s.CurrentHighestBid, s.CurrentHighestBidder),
		})
		if err != nil {
			log.Printf("Error sending auction update to client: %v", err)
			return err
		}
		time.Sleep(2 * time.Second)
	}
}
