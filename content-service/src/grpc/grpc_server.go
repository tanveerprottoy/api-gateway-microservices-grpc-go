package grpc

import (
	"fmt"
	"log"
	"net"
	"txp/contentservice/src/module/content"
	"txp/contentservice/src/module/content/proto"

	_grpc "google.golang.org/grpc"
)

var (
	Server *_grpc.Server
	lis    net.Listener
)

func InitServer() {
	var err error
	lis, err = net.Listen("tcp", fmt.Sprintf(":%d", 5001))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	Server = _grpc.NewServer()
}

func RegisterRPCs() {
	h := &content.ContentRPC{}
	h.InitDependencies()
	proto.RegisterContentServiceServer(
		Server,
		h,
	)
}

func Run() {
	log.Printf("server listening at %v", lis.Addr())
	if err := Server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
