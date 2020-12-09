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

	label := gui.NewLabel("Hello World!").
		WithWidgetProps(
			gui.WithAlign(gfx.AlignTop),
			gui.WithPadding(50),
		).
		WithTextProps(
			gui.WithFontSize(48),
		)

	app.NewScene().Add(label)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
