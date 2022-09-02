package user

import "net/http"

type UserHandler struct {
	service *UserService
}

func (h *UserHandler) InitDependencies() {
	// repo := &UserRepository{}
	h.service = &UserService{
		// repo: repo,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	h.service.Create(
		w,
		r,
	)
}

func (h *UserHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	h.service.FindAll(
		w,
		r,
	)
}

func (h *UserHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	h.service.FindOne(
		w,
		r,
	)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	h.service.Update(
		w,
		r,
	)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	h.service.Delete(
		w,
		r,
	)
}
