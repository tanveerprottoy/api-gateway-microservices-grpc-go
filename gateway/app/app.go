package app

import (
	"log"
	"net/http"
	"txp/gateway/app/module/user"

	"github.com/go-chi/chi"
)

// global var
var (
	UserModule *user.UserModule
)

// App struct
type App struct {
	router *Router
}

func (a *App) initModules() {
	UserModule = new(user.UserModule)
	UserModule.InitComponents()
}

// Init app
func (a *App) InitComponents() {
	a.initModules()
	a.router = NewRouter(
		chi.NewRouter(),
	)
}

// Run app
func (a *App) Run() {
	err := http.ListenAndServe(
		":8080",
		a.router.Mux,
	)
	if err != nil {
		log.Fatal(err)
	}
}
