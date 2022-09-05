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
	grpc.InitClientConn()
	grpc.InitServiceClients()
}

// Run app
func (a *App) Run() {
	log.Fatal(
		http.ListenAndServe(
			":8080",
			a.router.Mux,
		),
	)
}
