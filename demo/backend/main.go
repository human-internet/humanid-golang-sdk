package main

import (
	"PATH_TO/app"
)

func main() {
	app := &app.App{}
	app.Initialize()

	app.Run(":4000")
}
