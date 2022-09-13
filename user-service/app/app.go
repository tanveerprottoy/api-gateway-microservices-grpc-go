package app

import (
	"txp/userservice/app/module/user"
	"txp/userservice/pkg/grpc"
)

var (
	UserModule *user.UserModule
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
		UserModule.UserRPC,
	)
}

// Run app
func (a *App) Run() {
	a.grpcServer.Run()
}

func (a *App) initModules() {
	UserModule = new(user.UserModule)
	UserModule.InitComponents()
}
