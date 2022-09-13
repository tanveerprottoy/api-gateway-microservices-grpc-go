package main

import "txp/contentservice/src"

func main() {
	app := &src.App{}
	app.Init()
}

// Multiply just to check unit test
func Multiply() int {
	return 25 * 4
}
