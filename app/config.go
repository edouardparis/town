package app

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/LizardsTown/opennode"
	"github.com/pkg/errors"

	"github.com/EdouardParis/town/logging"
	"github.com/EdouardParis/town/store"
)

type Config struct {
	PaymentConfig opennode.Config `json:"opennode"`
	LoggerConfig  logging.Config  `json:"logger"`
	StoreConfig   store.Config    `json:"store"`
	InfoConfig    InfoConfig      `json:"info"`
}

type InfoConfig struct {
	URLs struct {
		Website string `json:"website"`
	} `json:"urls"`
}

func NewConfig(path string) (*Config, error) {
	if path != "" {
		return LoadConfigFile(path)
	}

	return NewConfigFromENV()
}

func LoadConfigFile(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer f.Close()

	var config Config
	err = json.NewDecoder(f).Decode(&config)
	return &config, errors.WithStack(err)
}

func NewConfigFromENV() (*Config, error) {
	var err error
	c := &Config{}
	c.PaymentConfig.APIKey = os.Getenv("OPENNODE_APIKEY")
	c.PaymentConfig.Debug = (os.Getenv("OPENNODE_DEBUG") == "true")

	c.LoggerConfig.Environment = os.Getenv("LOGGER_ENV")

	c.StoreConfig.Name = os.Getenv("POSTGRESQL_ADDON_DB")
	c.StoreConfig.Host = os.Getenv("POSTGRESQL_ADDON_HOST")
	c.StoreConfig.Password = os.Getenv("POSTGRESQL_ADDON_PASSWORD")
	c.StoreConfig.User = os.Getenv("POSTGRESQL_ADDON_USER")
	c.StoreConfig.SSLMode = os.Getenv("POSTGRESQL_ADDON_SSLMODE")
	c.StoreConfig.Port, err = strconv.Atoi(os.Getenv("POSTGRESQL_ADDON_PORT"))
	if err != nil {
		return nil, err
	}

	return c, nil
}
