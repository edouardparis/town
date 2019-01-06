package store

import (
	"context"

	lkb "github.com/ulule/loukoum/builder"
	"github.com/ulule/makroud"
)

// ----------------------------------------------------------------------------
// Store
// ----------------------------------------------------------------------------

type Store interface {
	Conn() makroud.Driver
	Get(context.Context, lkb.Builder, interface{}) error
	Find(context.Context, lkb.Builder, interface{}) error
	Exec(context.Context, lkb.Builder, ...interface{}) error
	Save(context.Context, lkb.Builder, interface{}) error
	Transaction(context.Context, func(Store) error) error
}
