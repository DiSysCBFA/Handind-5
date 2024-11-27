package Server

import (
	api "github.com/DiSysCBFA/Handind-5/Api"
)

type Server struct {
	api.UnimplementedAuctionserviceServer
	Port                 string
	CurrentHighestBid    int64
	CurrentHighestBidder string
}

func NewServer(port string) *Server {
	return &Server{
		Port: port,
	}
}

func SendBid(bid *api.Bid) (*api.BidAck, error) {
	return &api.BidAck{Ack: "Bid Accepted"}, nil
}

func JoinAuction(stream api.Auctionservice_JoinAuctionServer) error {
	return nil
}
