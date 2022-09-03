package user

import (
	"encoding/json"
	"log"
	"net/http"
	"txp/gateway/src/core"
	"txp/gateway/src/module/user/dto"
	"txp/gateway/src/module/user/proto"
	"txp/gateway/src/util"

	"github.com/go-chi/chi"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type UserService struct {
	client proto.UserServiceClient
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
	u, err := s.client.CreateUser(
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
	u, err := s.client.ReadUsers(
		r.Context(),
		nil,
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

func (s *UserService) ReadOne(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, core.UrlKeyId)
	u, err := s.client.ReadUser(
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
	userId := chi.URLParam(r, core.UrlKeyId)
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
	u, err := s.client.UpdateUser(
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
	userId := chi.URLParam(r, core.UrlKeyId)
	u, err := s.client.DeleteUser(
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
