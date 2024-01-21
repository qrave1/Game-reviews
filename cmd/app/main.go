package main

import (
	"gameReview/cmd/wire"
)

func main() {
	app, err := wire.InitializeContainer()
	if err != nil {
		return
	}

	err = app.Run()
	if err != nil {
		return
	}
}
