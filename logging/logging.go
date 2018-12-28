package logging

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zapcore.Field

type Logger interface {
	Info(string, ...Field)
}

func String(k, v string) Field {
	return zap.String(k, v)
}

func Duration(k string, d time.Duration) Field {
	return zap.Duration(k, d)
}

func Int(k string, i int) Field {
	return zap.Int(k, i)
}

func Error(v error) Field {
	return zap.Error(v)
}

type Config struct {
}

func NewCliLogger(c *Config) (Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{"stdout"}
	return cfg.Build()
}
