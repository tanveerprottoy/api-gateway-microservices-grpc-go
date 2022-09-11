package user

import (
	"net/http"
)

type UserHandler struct {
	service *UserService
}

func (h *UserHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	h.service.Create(
		w,
		r,
	)
}

func (h *UserHandler) ReadMany(
	w http.ResponseWriter,
	r *http.Request,
) {
	h.service.ReadMany(
		w,
		r,
	)
}

func (h *UserHandler) ReadOne(
	w http.ResponseWriter,
	r *http.Request,
) {
	h.service.ReadOne(
		w,
		r,
	)
}

func (h *UserHandler) Update(
	w http.ResponseWriter,
	r *http.Request,
) {
	h.service.Update(
		w,
		r,
	)
}

func (h *UserHandler) Delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	h.service.Delete(
		w,
		r,
	)
}
