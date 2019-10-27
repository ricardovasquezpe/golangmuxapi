package main

import "golangmuxapi/app"

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(":3000")
}
