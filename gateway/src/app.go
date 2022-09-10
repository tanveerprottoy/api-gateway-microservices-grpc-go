package src

import (
	"log"
	"net/http"
	"txp/gateway/src/grpc"
)

// App struct
type App struct {
	router *Router
}

// Init app
func (a *App) Init() {
	a.router = &Router{}
	a.router.Init()
	grpc.InitClientConns()
	grpc.InitServiceClients()
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
