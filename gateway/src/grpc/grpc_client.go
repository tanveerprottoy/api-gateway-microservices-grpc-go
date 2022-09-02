package grpc

import (
	"log"

	_grpc "google.golang.org/grpc"
)

var Conn *_grpc.ClientConn

func Init() {
	var err error
	Conn, err = _grpc.Dial(
		"0.0.0.0:5000",
		_grpc.WithInsecure(),
		_grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// defer g.conn.Close()
}
