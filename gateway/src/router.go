package src

import (
	"net/http"
	"txp/gateway/src/core/middleware"
	"txp/gateway/src/util"

	"github.com/go-chi/chi"
)

// Router struct
type Router struct {
	Mux  *chi.Mux
	TMux *http.ServeMux
}

func (r *Router) Init() {
	r.Mux = chi.NewRouter()
	r.registerMiddlewares()
	r.registerUserRoutes(
		util.V1,
	)
	/* r.registerContentRoutes(
		util.V1,
	) */
}

func (r *Router) registerMiddlewares() {
	r.Mux.Use(
		middleware.JSONContentTypeMiddleWare,
	)
}

func (r *Router) registerUserRoutes(
	version string,
) {
	r.Mux.Route(
		util.ApiPattern+version+util.UsersPattern,
		func(r chi.Router) {
			r.Get(
				util.RootPattern,
				UserModule.UserHandler.ReadMany,
			)
			r.Get(
				util.RootPattern+"{id}",
				UserModule.UserHandler.ReadOne,
			)
			r.Post(
				util.RootPattern,
				UserModule.UserHandler.Create,
			)
			r.Patch(
				util.RootPattern+"{id}",
				UserModule.UserHandler.Update,
			)
			r.Delete(
				util.RootPattern+"{id}",
				UserModule.UserHandler.Delete,
			)
		},
	)
}
