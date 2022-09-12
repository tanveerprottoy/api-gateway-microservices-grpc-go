package main

import "txp/gateway/app"

func main() {
	app := &app.App{}
	app.Init()
	app.Run()
}

// Multiply just to check unit test
func Multiply() int {
	return 25 * 4
}
