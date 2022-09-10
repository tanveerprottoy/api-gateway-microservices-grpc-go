package user

import (
	"context"
	"fmt"
	"log"
	"time"
	"txp/userservice/src/module/user/proto"
	"txp/userservice/src/util"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	if err != nil || l != "" {
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
	d := &proto.Users{}
	rows, err := s.repo.ReadMany()
	var (
		users      []*proto.User
		id         string
		name       string
		created_at time.Time
		updated_at time.Time
	)
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		if err := rows.Scan(&id, &name, &created_at, &updated_at); err != nil {
			return nil, fmt.Errorf("ReadMany %v", err)
		}
		users = append(users, &proto.User{
			Id:   id,
			Name: name,
			CreatedAt: timestamppb.New(
				created_at,
			),
			UpdatedAt: timestamppb.New(
				updated_at,
			),
		})
	}
	d.Users = users
	return d, err
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
	row := s.repo.ReadOne(
		strVal.Value,
	)
	if row == nil {
		return nil, util.RespondError(
			codes.NotFound,
			util.NotFound,
		)
	}
	var (
		id         string
		name       string
		created_at time.Time
		updated_at time.Time
	)
	if err := row.Scan(&id, &name, &created_at, &updated_at); err != nil {
		return nil, fmt.Errorf("ReadOne %v", err)
	}
	u := &proto.User{
		Id:   id,
		Name: name,
		CreatedAt: timestamppb.New(
			created_at,
		),
		UpdatedAt: timestamppb.New(
			updated_at,
		),
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
