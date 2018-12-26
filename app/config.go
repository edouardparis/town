package app

import (
	"git.iiens.net/edouardparis/town/logging"
	"git.iiens.net/edouardparis/town/store"
)

type Config struct {
	loggerConfig *logging.Config
	storeConfig  *store.Config
}
