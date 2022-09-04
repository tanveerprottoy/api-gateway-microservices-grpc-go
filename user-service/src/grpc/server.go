package grpc

import (
	"context"
	"log"
	"txp/userservice/src/module/user/proto"
)

type GServer struct{
	proto.UnimplementedUserServiceServer
}

func (g *GServer) ReadUsers(ctx context.Context, in *proto.VoidParam) (*proto.Users, error) {
	log.Printf("Received.ReadUsers:")
	return &proto.Users{}, nil
}
