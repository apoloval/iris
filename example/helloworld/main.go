package main

import (
	"github.com/apoloval/karen"
	"github.com/apoloval/karen/gfx"
	"github.com/apoloval/karen/gui"
)

func main() {
	app, err := karen.NewApp()
	if err != nil {
		panic(err)
	}

	label := gui.NewLabel("Hello World!")
	label.TextParams.Size = 48
	label.Align = gfx.AlignTop
	label.Padding = 50

	app.NewScene().Add(label)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
