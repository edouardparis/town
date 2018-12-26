package app

import (
	"git.iiens.net/edouardparis/town/logging"
	"git.iiens.net/edouardparis/town/store"
)

type App struct {
	Logger logging.Logger
	Store  store.Store
}

func New(c *Config) (*App, error) {
	logger, err := logging.NewCliLogger(&c.LoggerConfig)
	if err != nil {
		return nil, err
	}

	s, err := store.New(&c.StoreConfig, logger)
	if err != nil {
		return nil, err
	}

	return &App{Logger: logger, Store: s}, nil
}
