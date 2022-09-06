package user

import (
	"context"
	"log"
	"txp/userservice/src/module/user/proto"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type UserService struct {
	repo *UserRepository
}

func (s *UserService) Create(
	context.Context,
	*proto.User,
) (*proto.User, error) {
	return &proto.User{}, nil
}

func (s *UserService) ReadMany(
	ctx context.Context,
	v *proto.VoidParam,
) (*proto.Users, error) {
	log.Print("ReadMany rpc")
	return s.repo.FindMany(), nil
}

/* func (s *UserService) ReadUserStream(
	v *proto.VoidParam,
	serv proto.UserService_ReadUserStreamServer,
) (*proto.Users, error) {
	return &proto.Users{}, nil
} */

func (s *UserService) ReadOne(
	ctx context.Context,
	strVal *wrapperspb.StringValue,
) (*proto.User, error) {
	return &proto.User{}, nil
}

func (s *UserService) Update(
	ctx context.Context,
	u *proto.User,
) (*proto.User, error) {
	return &proto.User{}, nil
}

func (s *UserService) Delete(
	ctx context.Context,
	strVal *wrapperspb.StringValue,
) (*wrapperspb.BoolValue, error) {
	return &wrapperspb.BoolValue{Value: true}, nil
}
