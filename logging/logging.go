package logging

import (
	"go.uber.org/zap"
)

type Logger interface {
	Info(string, ...zap.Field)
}

func String(k, v string) zap.Field {
	return zap.String(k, v)
}

func Error(v error) zap.Field {
	return zap.Error(v)
}

type Config struct {
}

func NewCliLogger(c *Config) (Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{"stdout"}
	return cfg.Build()
}
