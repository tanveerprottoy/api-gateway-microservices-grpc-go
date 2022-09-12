package main

import "txp/userservice/app"

func main() {
	app := &app.App{}
	app.Init()
}

// Multiply just to check unit test
func Multiply() int {
	return 25 * 4
}
