package src

import (
	"txp/userservice/src/data"
	"txp/userservice/src/grpc"
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
