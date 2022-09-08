package src

import (
	"txp/contentservice/src/grpc"
)

// App struct
type App struct {
}

// Init app
func (a *App) Init() {
	grpc.InitServer()
	grpc.RegisterRPCs()
	grpc.Run()
}
