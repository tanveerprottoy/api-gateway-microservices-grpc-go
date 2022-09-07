package user

import (
	"context"
	"log"
	"txp/userservice/src/module/user/proto"
	"txp/userservice/src/util"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type UserService struct {
	repo *UserRepository
}

func (s *UserService) Create(
	ctx context.Context,
	u *proto.User,
) (*proto.User, error) {
	l, err := s.repo.Create(
		u,
	)
	if err != nil || l != nil {
		return nil, util.RespondError(
			codes.Unknown,
			util.UnknownError,
		)
	}
	return u, nil
}

func (s *UserService) ReadMany(
	ctx context.Context,
	v *proto.VoidParam,
) (*proto.Users, error) {
	log.Print("ReadMany rpc")
	return s.repo.ReadMany()
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
	u, err := s.repo.ReadOne(
		strVal.Value,
	)
	if err != nil {
		return nil, util.RespondError(
			codes.Unknown,
			util.UnknownError,
		)
	}
	if u != nil {
		return nil, util.RespondError(
			codes.NotFound,
			util.NotFound,
		)
	}
	return u, nil
}

func (s *UserService) Update(
	ctx context.Context,
	p *proto.UpdateUserParam,
) (*proto.User, error) {
	r, err := s.repo.Update(
		p.Id,
		p.User,
	)
	if err != nil || r <= 0 {
		return nil, util.RespondError(
			codes.Unknown,
			util.UnknownError,
		)
	}
	return p.User, nil
}

func (s *UserService) Delete(
	ctx context.Context,
	strVal *wrapperspb.StringValue,
) (*wrapperspb.BoolValue, error) {
	r, err := s.repo.Delete(
		strVal.Value,
	)
	if err != nil || r <= 0 {
		return nil, util.RespondError(
			codes.Unknown,
			util.UnknownError,
		)
	}
	return &wrapperspb.BoolValue{Value: true}, nil
}
