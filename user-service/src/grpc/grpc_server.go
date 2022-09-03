package grpc

import (
	"fmt"
	"log"
	"net"
	"txp/userservice/src/module/user"
	"txp/userservice/src/module/user/proto"

	_grpc "google.golang.org/grpc"
)

var Server *_grpc.Server

func InitServer() {
	var err error
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	Server = _grpc.NewServer()
	log.Printf("server listening at %v", lis.Addr())
	if err := Server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func RegisterRPCs() {
	h := &user.UserRPC{}
	h.InitDependencies()
	proto.RegisterUserServiceServer(
		Server,
		h.UnimplementedUserServiceServer,
	)
}
