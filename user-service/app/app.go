package app

import (
	"txp/userservice/pkg/data"
	"txp/userservice/pkg/grpc"
)

// App struct
type App struct {
}

// Init app
func (a *App) Init() {
	data.Init()
	grpc.InitServer()
	grpc.RegisterRPCs()
	grpc.Run()
}
