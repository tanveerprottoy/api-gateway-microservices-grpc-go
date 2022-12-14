package user

import (
	"encoding/json"
	"log"
	"net/http"
	"txp/gateway/app/module/user/dto"
	"txp/gateway/app/module/user/proto"
	"txp/gateway/pkg/grpc"
	"txp/gateway/pkg/util"

	"github.com/go-chi/chi"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type UserService struct {
}

func (s *UserService) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var b *dto.CreateUpdateUserDto
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		util.RespondError(http.StatusBadRequest, err, w)
		return
	}
	// ctx := context.Background()
	// send to service
	e, err := grpc.UserServiceClient.CreateUser(
		r.Context(),
		&proto.User{
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
	log.Print(e)
	util.Respond(
		http.StatusCreated,
		b,
		w,
	)
}

func (s *UserService) ReadMany(
	w http.ResponseWriter,
	r *http.Request,
) {
	u, err := grpc.UserServiceClient.ReadUsers(
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

func (s *UserService) ReadOne(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, util.UrlKeyId)
	u, err := grpc.UserServiceClient.ReadUser(
		r.Context(),
		&wrapperspb.StringValue{Value: userId},
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

func (s *UserService) Update(
	w http.ResponseWriter,
	r *http.Request,
) {
	userId := chi.URLParam(r, util.UrlKeyId)
	var b *dto.CreateUpdateUserDto
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		util.RespondError(
			http.StatusBadRequest,
			err,
			w,
		)
		return
	}
	u, err := grpc.UserServiceClient.UpdateUser(
		r.Context(),
		&proto.UpdateUserParam{
			Id: userId,
			User: &proto.User{
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

func (s *UserService) Delete(w http.ResponseWriter, r *http.Request) {
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
