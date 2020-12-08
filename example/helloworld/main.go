package main

import "github.com/apoloval/karen"

func main() {
	_, err := karen.NewApp()
	if err != nil {
		panic(err)
	}

}
