package store

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/ulule/loukoum/parser"

	"github.com/EdouardParis/town/logging"
)

// Option is a functional option.
type Option func(*Options)

// Options are redis wrapper options.
type Options struct {
	Logger     logging.Logger
	pghostname string
}

// NewOptions creates a new Options instance from given options.
func NewOptions(opts ...Option) Options {
	opt := Options{}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// WithLogger sets the logger instance.
func WithLogger(l logging.Logger) Option {
	return func(o *Options) {
		o.Logger = l
	}
}

// WithPGHostname sets the postgres hostname.
func WithPGHostname(pghostname string) Option {
	return func(o *Options) {
		o.pghostname = pghostname
	}
}

// Wrapper wrapps logging and satisfy makroud logger interface.
type Wrapper struct {
	hostname   string
	pghostname string
	logger     logging.Logger
}

// Log will emmit given queries on wrapper's attached Logger and Meter.
func (w *Wrapper) Log(query string, delta time.Duration) {
	analysis, err := parser.Analyze(query, parser.AnalyzerOption{
		Operation: true,
		Table:     true,
	})
	if err != nil {
		w.logger.Info("Cannot analyze query", logging.String("query", query), logging.Error(err))
		return
	}

	table := strings.Replace(analysis.Table, ".", "_", -1)
	operation := strings.Replace(analysis.Operation, ".", "_", -1)

	var metric string
	if w.pghostname == "" {
		metric = fmt.Sprintf("postgres.%s.%s.%s", w.hostname, table, operation)
	} else {
		metric = fmt.Sprintf("postgres.%s.%s.%s.%s", w.hostname, w.pghostname, table, operation)
	}
	w.logger.Info(metric, logging.Duration("duration", delta))
}

// NewWrapper creates a new wrapper with options.
func NewWrapper(options ...Option) (*Wrapper, error) {
	return NewWrapperWithOptions(NewOptions(options...))
}

// NewWrapperWithOptions creates a new wrapper from the given struct options.
func NewWrapperWithOptions(options Options) (*Wrapper, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, errors.Wrap(err, "cannot obtain server hostname")
	}

	return &Wrapper{
		hostname:   strings.Replace(hostname, ".", "_", -1),
		pghostname: strings.Replace(options.pghostname, ".", "_", -1),
		logger:     options.Logger,
	}, nil
}
