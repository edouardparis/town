package main

import (
	"context"
	"flag"

	"git.iiens.net/edouardparis/town/app"
	"git.iiens.net/edouardparis/town/web/server"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "c", "", "config file")
	flag.Parse()

	cfg, err := app.NewConfig(configFile)
	if err != nil {
		panic(err)
	}

	app, err := app.New(cfg)
	if err != nil {
		panic(err)
	}

	err = server.Run(context.Background(), app)
	app.Logger.Info(err.Error())
}
