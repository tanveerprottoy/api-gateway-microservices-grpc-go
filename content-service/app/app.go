package app

import (
	"txp/contentservice/app/module/content"
	"txp/contentservice/pkg/grpc"
)

var (
	ContentModule *content.ContentModule
)

// App struct
type App struct {
	grpcServer *grpc.GRPCServer
}

func NewApp() *App {
	a := new(App)
	a.grpcServer = grpc.NewGRPCServer()
	return a
}

// InitComponents app
func (a *App) InitComponents() {
	a.initModules()
	a.grpcServer.RegisterRPCs(
		ContentModule.ContentRPC,
	)
}

// Run app
func (a *App) Run() {
	a.grpcServer.Run()
}

func (a *App) initModules() {
	ContentModule = new(content.ContentModule)
	ContentModule.InitComponents()
}
