package app

import (
	"log"
	"net/http"
	"txp/gateway/app/grpc"
	"txp/gateway/app/module/user"
)

// global var for use
var (
	UserModule *user.UserModule
)

// App struct
type App struct {
	router *Router
}

// Init app
func (a *App) Init() {
	a.initModules()
	a.router = &Router{}
	a.router.Init()
	grpc.InitClientConns()
	grpc.InitServiceClients()
}

func (a *App) initModules() {
	UserModule = &user.UserModule{}
	UserModule.InitComponents()
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
