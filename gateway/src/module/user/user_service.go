package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"txp/gateway/src/core"
	"txp/gateway/src/module/user/dto"
	"txp/gateway/src/module/user/entity"
	"txp/gateway/src/util"

	"github.com/go-chi/chi"
)

type UserService struct {
	// repo *UserRepository
}

func (s *UserService) Create(w http.ResponseWriter, r *http.Request) {
	var b *dto.CreateUpdateUserBody
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		util.RespondError(http.StatusBadRequest, err, w)
		return
	}
	// send to service
	// here

	util.Respond(
		http.StatusCreated,
		map[string]bool{
			"created": true,
		},
		w,
	)
}

func (s *UserService) FindAll(w http.ResponseWriter, r *http.Request) {
	util.Respond(
		http.StatusOK,
		[]entity.User{},
		w,
	)
}

func (s *UserService) FindOne(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, core.UrlKeyId)
	_, err := strconv.Atoi(userId)
	if err != nil {
		util.RespondErrorMessage(
			http.StatusBadRequest,
			util.BadRequest,
			w,
		)
		return
	}
	util.Respond(
		http.StatusOK,
		entity.User{},
		w,
	)
}

func (s *UserService) Update(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, core.UrlKeyId)
	_, err := strconv.Atoi(userId)
	if err != nil {
		util.RespondErrorMessage(
			http.StatusBadRequest,
			util.BadRequest,
			w,
		)
		return
	}
	var b *dto.CreateUpdateUserBody
	err = json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		util.RespondError(http.StatusBadRequest, err, w)
		return
	}
	util.Respond(
		http.StatusOK,
		entity.User{},
		w,
	)
}

func (s *UserService) Delete(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, core.UrlKeyId)
	_, err := strconv.Atoi(userId)
	if err != nil {
		util.RespondErrorMessage(
			http.StatusBadRequest,
			util.BadRequest,
			w,
		)
		return
	}
	util.Respond(
		http.StatusOK,
		map[string]int64{util.RowsAffected: 1},
		w,
	)
}