package main

import "txp/gateway/src"

func main() {
	app := &src.App{}
	app.Init()
	app.Run()
}

// Multiply just to check unit test
func Multiply() int {
	return 25 * 4
}
