package main

import "github.com/apoloval/karen"

func main() {
	app, err := karen.NewApp()
	if err != nil {
		panic(err)
	}

	app.NewScene()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
