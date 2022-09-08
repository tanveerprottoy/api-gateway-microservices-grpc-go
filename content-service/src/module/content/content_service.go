package content

import (
	"context"
	"log"
	"txp/contentservice/src/module/content/proto"
	"txp/contentservice/src/util"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type ContentService struct {
	repo *ContentRepository
}

func (s *ContentService) Create(
	ctx context.Context,
	u *proto.Content,
) (*proto.Content, error) {
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

func (s *ContentService) ReadMany(
	ctx context.Context,
	v *proto.VoidParam,
) (*proto.Contents, error) {
	log.Print("ReadMany rpc")
	return s.repo.ReadMany()
}

/* func (s *ContentService) ReadContentStream(
	v *proto.VoidParam,
	serv proto.ContentService_ReadContentStreamServer,
) (*proto.Contents, error) {
	return &proto.Contents{}, nil
} */

func (s *ContentService) ReadOne(
	ctx context.Context,
	strVal *wrapperspb.StringValue,
) (*proto.Content, error) {
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

func (s *ContentService) Update(
	ctx context.Context,
	p *proto.UpdateContentParam,
) (*proto.Content, error) {
	r, err := s.repo.Update(
		p.Id,
		p.Content,
	)
	if err != nil || r <= 0 {
		return nil, util.RespondError(
			codes.Unknown,
			util.UnknownError,
		)
	}
	return p.Content, nil
}

func (s *ContentService) Delete(
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
