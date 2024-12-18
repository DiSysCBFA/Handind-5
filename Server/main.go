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
	HighestBid *api.Bid
	Timestamp  int64
}

func (s *Server) TryBid(ctx context.Context, incommingBid *api.Bid) (*api.BidAck, error) {
	if incommingBid.Timestamp > (s.Timestamp + 15000000000) {
		bidAck := &api.BidAck{
			Ack:       "Auction Ended",
			Timestamp: time.Now().UnixNano(),
		}

		return bidAck, nil
	}
	if s.HighestBid == nil || incommingBid.Bid > s.HighestBid.Bid {
		s.HighestBid = incommingBid
		bidAck := &api.BidAck{
			Ack:       "accepted",
			Timestamp: time.Now().UnixNano(),
		}

		return bidAck, nil
	} else {
		bidAck := &api.BidAck{
			Ack:       "rejected",
			Timestamp: time.Now().UnixNano(),
		}

		return bidAck, nil
	}
}

func StartServer() *Server {

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	server := &Server{
		Timestamp: time.Now().UnixNano(),
	}

	// Register the pool with the gRPC server
	api.RegisterAuctionserviceServer(grpcServer, server)

	// Create a TCP listener at port 5101
	for _, port := range []string{":4000", ":4001", ":4002"} {
		listener, err := net.Listen("tcp", port)
		if err != nil {
			log.Println("Error creating the server %v", err)
			continue
		} else {
			log.Println("Server started at port", port)

			if err := grpcServer.Serve(listener); err != nil {
				log.Fatalf("Error creating the server %v", err)
			}
			break
		}

	}
	return server
}
func (s *Server) JoinAuction(ctx context.Context, empty *api.Empty) (*api.AuctionResult, error) {
	result := &api.AuctionResult{
		Status: "Highest Bid:" + fmt.Sprint(s.HighestBid.Bid) + " by: " + s.HighestBid.Bidder,
	}
	return result, nil
}
