package user

import (
	"encoding/json"
	"log"
	"net/http"
	"txp/gateway/src/grpc"
	"txp/gateway/src/module/user/dto"
	"txp/gateway/src/module/user/proto"
	"txp/gateway/src/util"

	"github.com/go-chi/chi"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type UserService struct {
}

func (s *UserService) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	var b *dto.CreateUpdateUserBody
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		util.RespondError(http.StatusBadRequest, err, w)
		return
	}
	// ctx := context.Background()
	// send to service
	u, err := grpc.UserClient.CreateUser(
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
	log.Print(u)
	util.Respond(
		http.StatusCreated,
		map[string]bool{
			"created": true,
		},
		w,
	)
}

func (s *UserService) ReadMany(
	w http.ResponseWriter,
	r *http.Request,
) {
	u, err := grpc.UserClient.ReadUsers(
		r.Context(),
		&proto.VoidParam{},
	)
/* 	c := proto.NewUserServiceClient(grpc.Conn)
	res, err := c.ReadUsers(r.Context(), &proto.VoidParam{})
	log.Print(res) */
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
	u, err := grpc.UserClient.ReadUser(
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

func (s *UserService) Update(
	w http.ResponseWriter,
	r *http.Request,
) {
	userId := chi.URLParam(r, util.UrlKeyId)
	var b *dto.CreateUpdateUserBody
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		util.RespondError(
			http.StatusBadRequest,
			err,
			w,
		)
		return
	}
	u, err := grpc.UserClient.UpdateUser(
		r.Context(),
		&proto.User{
			Id:   userId,
			Name: b.Name,
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
	u, err := grpc.UserClient.DeleteUser(
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
