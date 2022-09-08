package grpc

import (
	"log"
	_contentProto "txp/gateway/src/module/content/proto"
	"txp/gateway/src/module/user/proto"

	_grpc "google.golang.org/grpc"
)

var (
	Conns                [2]*_grpc.ClientConn
	UserServiceClient    proto.UserServiceClient
	ContentServiceClient _contentProto.ContentServiceClient
)

func InitClientConns() {
	var err error
	Conns[0], err = _grpc.Dial(
		"0.0.0.0:5000",
		_grpc.WithInsecure(),
		_grpc.WithBlock(),
	)
	if err != nil {
		log.Printf("userservice did not connect: %v", err)
	}
	Conns[1], err = _grpc.Dial(
		"0.0.0.0:5001",
		_grpc.WithInsecure(),
		_grpc.WithBlock(),
	)
	if err != nil {
		log.Printf("contentservice did not connect: %v", err)
	}
}

func InitServiceClients() {
	UserServiceClient = proto.NewUserServiceClient(
		Conns[0],
	)
	ContentServiceClient = _contentProto.NewContentServiceClient(
		Conns[1],
	)
}
