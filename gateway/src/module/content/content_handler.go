package content

type ContentHandler struct {
	service *ContentService
}

/* func (h *ContentHandler) InitDependencies() {
	repo := &ContentRepository{}
	h.service = &ContentService{
		repo: repo,
	}
}

func (h *ContentHandler) Create(w http.ResponseWriter, r *http.Request) {
	h.service.Create(
		w,
		r,
	)
}

func (h *ContentHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	h.service.FindAll(
		w,
		r,
	)
}

func (h *ContentHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	h.service.FindOne(
		w,
		r,
	)
}

func (h *ContentHandler) Update(w http.ResponseWriter, r *http.Request) {
	h.service.Update(
		w,
		r,
	)
}

func (h *ContentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	h.service.Delete(
		w,
		r,
	)
} */
