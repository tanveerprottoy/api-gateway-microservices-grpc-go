package grpc

import (
	"log"
	"txp/gateway/src/module/user/proto"

	_grpc "google.golang.org/grpc"
)

var (
	Conn       *_grpc.ClientConn
	UserClient proto.UserServiceClient
)

func InitClientConn() {
	var err error
	Conn, err = _grpc.Dial(
		"0.0.0.0:5000",
		_grpc.WithInsecure(),
		_grpc.WithBlock(),
	)
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
}

func InitServiceClients() {
	UserClient = proto.NewUserServiceClient(
		Conn,
	)
}
