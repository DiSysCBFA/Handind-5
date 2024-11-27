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
