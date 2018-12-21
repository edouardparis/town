package app

import (
	"git.iiens.net/edouardparis/town/logging"
)

type App struct {
	Logger logging.Logger
}

func New() (*App, error) {
	logger, err := logging.NewCliLogger(&logging.Config{})
	if err != nil {
		return nil, err
	}

	return &App{logger}, nil
}
