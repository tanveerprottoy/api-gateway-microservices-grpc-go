package user

import (
	"context"
	"txp/userservice/src/module/user/proto"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type UserRPC struct {
	proto.UnimplementedUserServiceServer
	service *UserService
}

func (h *UserRPC) InitDependencies() {
	h.service = &UserService{}
}

func (h *UserRPC) CreateUser(
	ctx context.Context,
	u *proto.User,
) (*proto.User, error) {
	return h.service.Create(
		ctx,
		u,
	)
}

func (h *UserRPC) ReadUsers(
	ctx context.Context,
	v *proto.VoidParam,
) (*proto.Users, error) {
	return h.service.ReadMany(
		ctx,
		v,
	)
}

func (h *UserRPC) ReadUserStream(
	v *proto.VoidParam,
	serv proto.UserService_ReadUserStreamServer,
) (*proto.Users, error) {
	return nil, nil
	/* h.service.ReadMany(
		ctx,
		v,
	) */
}

func (h *UserRPC) ReadUser(
	ctx context.Context,
	strVal *wrapperspb.StringValue,
) (*proto.User, error) {
	return h.service.ReadOne(
		ctx,
		strVal,
	)
}

func (h *UserRPC) UpdateUser(
	ctx context.Context,
	u *proto.User,
) (*proto.User, error) {
	return h.service.Update(
		ctx,
		u,
	)
}

func (h *UserRPC) DeleteUser(
	ctx context.Context,
	strVal *wrapperspb.StringValue,
) (*wrapperspb.BoolValue, error) {
	return h.service.Delete(
		ctx,
		strVal,
	)
}