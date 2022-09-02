package content

type ContentService struct {
	// repo *ContentRepository
}

/* func (s *ContentService) Create(w http.ResponseWriter, r *http.Request) {
	var p *dto.CreateUpdateContentBody
	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		util.ErrorAbort(
			http.StatusBadRequest,
			util.BadRequest,
			w,
		r,
		)
		return
	}
	_, err = s.repo.Create(
		p,
	)
	if err != nil {
		util.ErrorAbort(
			http.StatusInternalServerError,
			err.Error(),
			w,
		r,
		)
		return
	}
	util.Respond(
		http.StatusCreated,
		map[string]bool{
			"created": true,
		},
		w,
		r,
	)
}

func (s *ContentService) FindAll(w http.ResponseWriter, r *http.Request) {
	Contents, err := s.repo.FindAll()
	if err != nil {
		util.ErrorAbort(
			http.StatusBadRequest,
			util.BadRequest,
			w,
		r,
		)
		return
	}
	util.Respond(
		http.StatusOK,
		Contents,
		w,
		r,
	)
}

func (s *ContentService) FindOne(w http.ResponseWriter, r *http.Request) {
	id, exists := ctx.Params.Get("id")
	if !exists {
		util.ErrorAbort(
			http.StatusBadRequest,
			util.BadRequest,
			w,
		r,
		)
		return
	}
	Content, err := s.repo.FindOne(
		id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			util.ErrorAbort(
				http.StatusNotFound,
				util.NotFound,
				w,
		r,
			)
			return
		}
		ctx.AbortWithError(
			http.StatusInternalServerError,
			errors.New(
				util.InternalServerError,
			),
		)
		/* util.ErrorAbort(
			http.StatusInternalServerError,
			util.InternalServerError,
			w,
		r,
		)
		return
	}
	util.Respond(
		http.StatusOK,
		Content,
		w,
		r,
	)
}

func (s *ContentService) Update(w http.ResponseWriter, r *http.Request) {
	id, exists := ctx.Params.Get("id")
	if !exists {
		util.ErrorAbort(
			http.StatusBadRequest,
			util.BadRequest,
			w,
		r,
		)
		return
	}
	var p *entity.Content
	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		util.ErrorAbort(
			http.StatusBadRequest,
			util.BadRequest,
			w,
		r,
		)
		return
	}
	rows, err := s.repo.Update(
		id,
		p,
	)
	if err != nil {
		util.ErrorAbort(
			http.StatusInternalServerError,
			util.InternalServerError,
			w,
		r,
		)
		return
	}
	if rows > 0 {
		util.Respond(
			http.StatusOK,
			map[string]int64{util.RowsAffected: rows},
			w,
		r,
		)
		return
	}
	util.ErrorAbort(
		http.StatusInternalServerError,
		util.InternalServerError,
		w,
		r,
	)
}

func (s *ContentService) Delete(w http.ResponseWriter, r *http.Request) {
	id, exists := ctx.Params.Get("id")
	if !exists {
		util.RespondError(
			http.StatusBadRequest,
			errors.New(
				util.BadRequest,
			),
			w,
		r,
		)
		return
	}
	rows, err := s.repo.Delete(
		id,
	)
	if err != nil {
		util.ErrorAbort(
			http.StatusInternalServerError,
			util.InternalServerError,
			w,
		r,
		)
		return
	}
	util.Respond(
		http.StatusOK,
		map[string]int64{util.RowsAffected: rows},
		w,
		r,
	)
} */
