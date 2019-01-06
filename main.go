package main

import (
	"context"
	"flag"

	"github.com/EdouardParis/town/app"
	"github.com/EdouardParis/town/web/server"
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
