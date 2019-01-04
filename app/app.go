package app

import (
	"github.com/EdouardParis/town/logging"
	"github.com/EdouardParis/town/models"
	"github.com/EdouardParis/town/store"
)

type App struct {
	Logger logging.Logger
	Store  store.Store
	Info   *models.Info
	Config *Config
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

	return &App{
		Logger: logger,
		Store:  s,
		Info:   NewInfo(&c.InfoConfig),
		Config: c,
	}, nil
}

func NewInfo(c *InfoConfig) *models.Info {
	return &models.Info{
		URLs: models.URLs{
			Website: c.URLs.Website,
		},
	}
}
