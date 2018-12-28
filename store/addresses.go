package store

import (
	"context"

	"github.com/pkg/errors"
	lk "github.com/ulule/loukoum"

	"git.iiens.net/edouardparis/town/models"
)

type Addresses struct {
	Store
}

func NewAddresses(s Store) *Addresses {
	return &Addresses{s}
}

func (a *Addresses) Create(ctx context.Context, address *models.Address) error {
	query := lk.Insert(address.TableName()).
		Set(
			lk.Pair("value", address.Value),
			lk.Pair("amount_collected", address.AmountCollected),
			lk.Pair("amount_received", address.AmountReceived),
			lk.Pair("created_at", address.CreatedAt),
			lk.Pair("updated_at", address.UpdatedAt),
		).Returning("id")
	err := a.Save(ctx, query, address)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
