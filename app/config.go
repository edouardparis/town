package app

import (
	"encoding/json"
	"net"
	"net/url"
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
	c.InfoConfig.URLs.Website = os.Getenv("INFO_URLS_WEBSITE")
	c.PaymentConfig.APIKey = os.Getenv("OPENNODE_APIKEY")
	c.PaymentConfig.Debug = (os.Getenv("OPENNODE_DEBUG") == "true")

	c.LoggerConfig.Environment = os.Getenv("LOGGER_ENV")

	c.StoreConfig, err = NewDBConfig()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func NewDBConfig() (store.Config, error) {
	c := store.Config{}
	s := os.Getenv("DATABASE_URL")
	if s != "" {
		u, err := url.Parse(s)
		if err != nil {
			return c, err
		}
		c.Name = u.Path
		c.User = u.User.Username()
		c.Password, _ = u.User.Password()
		host, port, _ := net.SplitHostPort(u.Host)
		c.Host = host

		c.Port, err = strconv.Atoi(port)
		if err != nil {
			return c, err
		}
		c.SSLMode = "require"
	}

	var err error
	c.Name = os.Getenv("POSTGRESQL_ADDON_DB")
	c.Host = os.Getenv("POSTGRESQL_ADDON_HOST")
	c.Password = os.Getenv("POSTGRESQL_ADDON_PASSWORD")
	c.User = os.Getenv("POSTGRESQL_ADDON_USER")
	c.SSLMode = os.Getenv("POSTGRESQL_ADDON_SSLMODE")
	c.Port, err = strconv.Atoi(os.Getenv("POSTGRESQL_ADDON_PORT"))
	return c, err
}
