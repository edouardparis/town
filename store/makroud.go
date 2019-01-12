package store

import (
	"context"

	"github.com/pkg/errors"
	lkb "github.com/ulule/loukoum/builder"
	"github.com/ulule/makroud"

	"github.com/EdouardParis/town/logging"
	"github.com/EdouardParis/town/models"
)

// ----------------------------------------------------------------------------
// connector
// ----------------------------------------------------------------------------

func NewConnector(cfg *Config, logger logging.Logger) (makroud.Driver, error) {
	wrapper, err := NewWrapper(
		WithPGHostname(cfg.Host),
		WithLogger(logger),
	)
	if err != nil {
		return nil, err
	}

	dbx, err := makroud.New(
		makroud.Host(cfg.Host),
		makroud.Port(cfg.Port),
		makroud.User(cfg.User),
		makroud.Password(cfg.Password),
		makroud.Database(cfg.Name),
		makroud.EnableSavepoint(),
		makroud.WithLogger(wrapper),
	)
	if err != nil {
		return nil, err
	}

	err = dbx.Ping()
	if err != nil {
		return nil, err
	}

	return dbx, nil
}

type store struct {
	db makroud.Driver
}

func (s *store) Conn() makroud.Driver {
	return s.db
}

// Get is a shortcut to retrieve a slice of instances from a Loukoum builder.
func (s *store) Get(ctx context.Context, builder lkb.Builder, dest interface{}) error {
	return makroud.Exec(ctx, s.db, builder, dest)
}

// Find is a shortcut to retrieve a slice of instances from a Loukoum builder.
func (s *store) Find(ctx context.Context, builder lkb.Builder, dest interface{}) error {
	err := makroud.Exec(ctx, s.db, builder, dest)
	if err != nil && !makroud.IsErrNoRows(err) {
		return err
	}
	return nil
}

// Save is a shortcut to create / update from a Loukoum builder.
// It will also mutate the given object to match the row values.
func (s *store) Save(ctx context.Context, builder lkb.Builder, dest interface{}) error {
	return s.Exec(ctx, builder, dest)
}

// Exec will execute given query.
// It will also mutate the given object to match the row values.
func (s *store) Exec(ctx context.Context, builder lkb.Builder, dest ...interface{}) error {
	return makroud.Exec(ctx, s.db, builder, dest...)
}

// Transaction executes f in a SQL transaction.
func (s *store) Transaction(ctx context.Context, f func(stx Store) error) error {
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "can't begin transaction")
	}

	err = f(&store{tx})
	if err != nil {
		rerr := tx.Rollback()
		if rerr != nil {
			//tracer.FromContext(ctx).Capture(rerr, nil)
		}
		if err == ErrDontCommit {
			return nil
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "can't commit transaction")
	}

	return nil
}

var ErrDontCommit = errors.New("don't commit transaction and don't return an error")

func New(c *Config, logger logging.Logger) (Store, error) {
	s, err := NewConnector(c, logger)
	if err != nil {
		return nil, err
	}

	setCachedColumns(s)

	return &store{s}, nil
}

// ----------------------------------------------------------------------------
// columns
// ----------------------------------------------------------------------------

var cachedColumns = make(map[string]string)

func setCachedColumns(driver makroud.Driver) {
	var err error
	for _, model := range []makroud.Model{
		models.Article{},
		models.Address{},
		models.Node{},
		models.Slug{},
		models.Order{},
		models.Reaction{},
		models.Comment{},
	} {
		cachedColumns[model.TableName()], err = getColumns(driver, model)
		if err != nil {
			panic(err)
		}
	}
}

func columns(model makroud.Model) string {
	return cachedColumns[model.TableName()]
}

// GetColumns returns a comma-separated string representation of a model's table columns.
func getColumns(driver makroud.Driver, model makroud.Model) (string, error) {
	schema, err := makroud.GetSchema(driver, model)
	if err != nil {
		return "", errors.Wrap(err, "sqlx: cannot fetch schema informations")
	}

	columns := schema.ColumnPaths().String()

	return columns, nil
}
