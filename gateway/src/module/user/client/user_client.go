package client

import (
	context "context"
	"txp/gateway/src/module/user/proto"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type UserClient struct {
}

func (c *UserClient) CreateUser(ctx context.Context, in *proto.User, opts ...grpc.CallOption) (*proto.User, error) {
	return nil, nil
}

func (c *UserClient) ReadUsers(ctx context.Context, in *proto.VoidParam, opts ...grpc.CallOption) (*proto.Users, error) {
	return nil, nil
}

func (c *UserClient) ReadUserStream(ctx context.Context, in *proto.VoidParam, opts ...grpc.CallOption) (proto.UserService_ReadUserStreamClient, error) {
	return nil, nil
}

func (c *UserClient) ReadUser(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*proto.User, error) {
	return nil, nil
}

func (c *UserClient) UpdateUser(ctx context.Context, in *proto.User, opts ...grpc.CallOption) (*proto.User, error) {
	return nil, nil
}

func (c *UserClient) DeleteUser(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	return nil, nil
}
