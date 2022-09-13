package grpc

import (
	"fmt"
	"log"
	"net"
	"txp/contentservice/app/module/content"
	"txp/contentservice/app/module/content/proto"

	_grpc "google.golang.org/grpc"
)

type GRPCServer struct {
	Server *_grpc.Server
	lis    net.Listener
}

func NewGRPCServer() *GRPCServer {
	var err error
	g := new(GRPCServer)
	g.lis, err = net.Listen("tcp", fmt.Sprintf(":%d", 5001))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	g.Server = _grpc.NewServer()
	return g
}

func (g *GRPCServer) RegisterRPCs(
	r *content.ContentRPC,
) {
	proto.RegisterContentServiceServer(
		g.Server,
		r,
	)
}

func (g *GRPCServer) Run() {
	log.Printf("server listening at %v", g.lis.Addr())
	if err := g.Server.Serve(g.lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
