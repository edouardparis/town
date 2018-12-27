package app

import (
	"git.iiens.net/edouardparis/town/logging"
	"git.iiens.net/edouardparis/town/models"
	"git.iiens.net/edouardparis/town/store"
)

type App struct {
	Logger logging.Logger
	Store  store.Store
	Info   *models.Info
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

	i := NewInfo(&c.InfoConfig)

	return &App{Logger: logger, Store: s, Info: i}, nil
}

func NewInfo(c *InfoConfig) *models.Info {
	return &models.Info{
		URLs: models.URLs{
			Website: c.URLs.Website,
		},
	}
}
