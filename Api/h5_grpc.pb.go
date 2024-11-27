// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: Api/h5.proto

package Api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Auctionservice_SendBid_FullMethodName     = "/Auction_service.Auctionservice/SendBid"
	Auctionservice_JoinAuction_FullMethodName = "/Auction_service.Auctionservice/JoinAuction"
)

// AuctionserviceClient is the client API for Auctionservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuctionserviceClient interface {
	SendBid(ctx context.Context, in *Bid, opts ...grpc.CallOption) (*BidAck, error)
	JoinAuction(ctx context.Context, in *Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Auction], error)
}

type auctionserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuctionserviceClient(cc grpc.ClientConnInterface) AuctionserviceClient {
	return &auctionserviceClient{cc}
}

func (c *auctionserviceClient) SendBid(ctx context.Context, in *Bid, opts ...grpc.CallOption) (*BidAck, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BidAck)
	err := c.cc.Invoke(ctx, Auctionservice_SendBid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionserviceClient) JoinAuction(ctx context.Context, in *Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Auction], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Auctionservice_ServiceDesc.Streams[0], Auctionservice_JoinAuction_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Empty, Auction]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Auctionservice_JoinAuctionClient = grpc.ServerStreamingClient[Auction]

// AuctionserviceServer is the server API for Auctionservice service.
// All implementations must embed UnimplementedAuctionserviceServer
// for forward compatibility.
type AuctionserviceServer interface {
	SendBid(context.Context, *Bid) (*BidAck, error)
	JoinAuction(*Empty, grpc.ServerStreamingServer[Auction]) error
	mustEmbedUnimplementedAuctionserviceServer()
}

// UnimplementedAuctionserviceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuctionserviceServer struct{}

func (UnimplementedAuctionserviceServer) SendBid(context.Context, *Bid) (*BidAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBid not implemented")
}
func (UnimplementedAuctionserviceServer) JoinAuction(*Empty, grpc.ServerStreamingServer[Auction]) error {
	return status.Errorf(codes.Unimplemented, "method JoinAuction not implemented")
}
func (UnimplementedAuctionserviceServer) mustEmbedUnimplementedAuctionserviceServer() {}
func (UnimplementedAuctionserviceServer) testEmbeddedByValue()                        {}

// UnsafeAuctionserviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuctionserviceServer will
// result in compilation errors.
type UnsafeAuctionserviceServer interface {
	mustEmbedUnimplementedAuctionserviceServer()
}

func RegisterAuctionserviceServer(s grpc.ServiceRegistrar, srv AuctionserviceServer) {
	// If the following call pancis, it indicates UnimplementedAuctionserviceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Auctionservice_ServiceDesc, srv)
}

func _Auctionservice_SendBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionserviceServer).SendBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auctionservice_SendBid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionserviceServer).SendBid(ctx, req.(*Bid))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auctionservice_JoinAuction_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AuctionserviceServer).JoinAuction(m, &grpc.GenericServerStream[Empty, Auction]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Auctionservice_JoinAuctionServer = grpc.ServerStreamingServer[Auction]

// Auctionservice_ServiceDesc is the grpc.ServiceDesc for Auctionservice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auctionservice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Auction_service.Auctionservice",
	HandlerType: (*AuctionserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendBid",
			Handler:    _Auctionservice_SendBid_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinAuction",
			Handler:       _Auctionservice_JoinAuction_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "Api/h5.proto",
}