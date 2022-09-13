package app

import (
	"log"
	"net/http"
	"txp/gateway/app/module/content"
	"txp/gateway/app/module/user"
)

// global var
var (
	UserModule    *user.UserModule
	ContentModule *content.ContentModule
)

// App struct
type App struct {
	router *Router
}

func (a *App) initModules() {
	UserModule = new(user.UserModule)
	UserModule.InitComponents()
	ContentModule = new(content.ContentModule)
	ContentModule.InitComponents()
}

// Init app
func (a *App) InitComponents() {
	a.initModules()
	a.router = NewRouter()
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
