package Server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	pb "github.com/DiSysCBFA/Handind-5/Api" // Import generated proto package
	"google.golang.org/grpc"
)

// AuctionServer represents the gRPC server
type AuctionServer struct {
	pb.UnimplementedAuctionServiceServer
	mu            sync.Mutex
	highestBid    int32
	highestBidder string
	isClosed      bool
	startTime     time.Time
	duration      time.Duration
}

// NewAuctionServer creates a new AuctionServer
func NewAuctionServer(duration time.Duration) *AuctionServer {
	return &AuctionServer{
		highestBid:    0,
		highestBidder: "",
		isClosed:      false,
		startTime:     time.Now(),
		duration:      duration,
	}
}

// Bid handles bid requests
func (s *AuctionServer) Bid(ctx context.Context, req *pb.BidRequest) (*pb.BidResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isClosed {
		return &pb.BidResponse{Status: pb.BidResponse_FAIL}, nil
	}

	if req.Amount > s.highestBid {
		s.highestBid = req.Amount
		s.highestBidder = req.Bidder
		return &pb.BidResponse{Status: pb.BidResponse_SUCCESS}, nil
	}

	return &pb.BidResponse{Status: pb.BidResponse_FAIL}, nil
}

// Result handles result requests
func (s *AuctionServer) Result(ctx context.Context, req *pb.ResultRequest) (*pb.ResultResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.isClosed && time.Since(s.startTime) > s.duration {
		s.isClosed = true
	}

	return &pb.ResultResponse{
		HighestBidder:   s.highestBidder,
		HighestBid:      s.highestBid,
		IsAuctionClosed: s.isClosed,
	}, nil
}

func main() {
	// Start gRPC server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	auctionServer := NewAuctionServer(100 * time.Second)

	pb.RegisterAuctionServiceServer(grpcServer, auctionServer)

	log.Println("Auction server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
