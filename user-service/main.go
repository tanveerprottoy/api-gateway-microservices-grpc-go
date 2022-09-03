package main

import "txp/userservice/src"

func main() {
	app := &src.App{}
	app.Init()
}

// Multiply just to check unit test
func Multiply() int {
	return 25 * 4
}
