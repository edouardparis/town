package main

import (
	"context"

	"git.iiens.net/edouardparis/town/app"
	"git.iiens.net/edouardparis/town/web/server"
)

func main() {
	app, err := app.New(&app.Config{})
	if err != nil {
		panic(err)
	}

	err = server.Run(context.Background(), app)
	app.Logger.Info(err.Error())
}
