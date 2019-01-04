package app

import (
	"encoding/json"
	"os"

	"github.com/LizardsTown/opennode"
	"github.com/pkg/errors"

	"git.iiens.net/edouardparis/town/logging"
	"git.iiens.net/edouardparis/town/store"
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

	return &Config{}, nil
}
