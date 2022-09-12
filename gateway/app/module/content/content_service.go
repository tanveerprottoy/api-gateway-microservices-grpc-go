package content

import (
	"encoding/json"
	"net/http"
	"txp/gateway/app/module/content/dto"
	"txp/gateway/app/module/content/proto"
	"txp/gateway/pkg/grpc"
	"txp/gateway/pkg/util"

	"github.com/go-chi/chi"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type ContentService struct {
}

func (s *ContentService) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var b *dto.CreateUpdateContentDto
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		util.RespondError(http.StatusBadRequest, err, w)
		return
	}
	// ctx := context.Background()
	// send to service
	_, err = grpc.ContentServiceClient.CreateContent(
		r.Context(),
		&proto.Content{
			Name: b.Name,
		},
	)
	if err != nil {
		util.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	util.Respond(
		http.StatusCreated,
		b,
		w,
	)
}

func (s *ContentService) ReadMany(
	w http.ResponseWriter,
	r *http.Request,
) {
	u, err := grpc.ContentServiceClient.ReadContents(
		r.Context(),
		&proto.VoidParam{},
	)
	if err != nil {
		util.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	util.Respond(
		http.StatusOK,
		u,
		w,
	)
}

func (s *ContentService) ReadOne(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, util.UrlKeyId)
	u, err := grpc.ContentServiceClient.ReadContent(
		r.Context(),
		&wrapperspb.StringValue{Value: userId},
		nil,
	)
	if err != nil {
		util.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	util.Respond(
		http.StatusOK,
		u,
		w,
	)
}

func (s *ContentService) Update(
	w http.ResponseWriter,
	r *http.Request,
) {
	userId := chi.URLParam(r, util.UrlKeyId)
	var b *dto.CreateUpdateContentDto
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		util.RespondError(
			http.StatusBadRequest,
			err,
			w,
		)
		return
	}
	u, err := grpc.ContentServiceClient.UpdateContent(
		r.Context(),
		&proto.UpdateContentParam{
			Id: userId,
			Content: &proto.Content{
				Name: b.Name,
			},
		},
		nil,
	)
	if err != nil {
		util.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	util.Respond(
		http.StatusOK,
		u,
		w,
	)
}

func (s *ContentService) Delete(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, util.UrlKeyId)
	u, err := grpc.UserServiceClient.DeleteUser(
		r.Context(),
		&wrapperspb.StringValue{Value: userId},
		nil,
	)
	if err != nil {
		util.RespondError(
			http.StatusInternalServerError,
			err,
			w,
		)
		return
	}
	util.Respond(
		http.StatusOK,
		u,
		w,
	)
}
