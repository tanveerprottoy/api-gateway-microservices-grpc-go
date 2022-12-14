package grpc

import (
	"fmt"
	"log"
	"net"
	"txp/userservice/app/module/user"
	"txp/userservice/app/module/user/proto"

	_grpc "google.golang.org/grpc"
)

type GRPCServer struct {
	Server *_grpc.Server
	lis    net.Listener
}

func NewGRPCServer() *GRPCServer {
	var err error
	g := new(GRPCServer)
	g.lis, err = net.Listen("tcp", fmt.Sprintf(":%d", 5000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	g.Server = _grpc.NewServer()
	return g
}

func (g *GRPCServer) RegisterRPCs(
	r *user.UserRPC,
) {
	proto.RegisterUserServiceServer(
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
