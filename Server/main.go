package Server

import (
	"context"
	"time"

	api "github.com/DiSysCBFA/Handind-5/Api"
)

type Server struct {
	api.UnimplementedAuctionserviceServer
	HighestBid *api.Bid
	Timestamp  int64
}

func (s *Server) Bid(ctx context.Context, incommingBid *api.Bid) (*api.BidAck, error) {
	if incommingBid.Timestamp > s.Timestamp+15000000000 {
		bidAck := &api.BidAck{
			Ack:       "late",
			Timestamp: time.Now().UnixNano(),
		}

		return bidAck, nil
	}
	if incommingBid.Bid > s.HighestBid.Bid || s.HighestBid == nil {
		s.HighestBid = incommingBid
	}

}
