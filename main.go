package main

import "muxgoapi/app"

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(":3000")
}
