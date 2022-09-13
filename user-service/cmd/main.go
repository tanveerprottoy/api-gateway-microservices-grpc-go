package main

import "txp/userservice/app"

func main() {
	app := app.NewApp()
	app.InitComponents()
	app.Run()
}

// Multiply just to check unit test
func Multiply() int {
	return 25 * 4
}
