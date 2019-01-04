package store

import (
	"context"

	"github.com/pkg/errors"
	lk "github.com/ulule/loukoum"
	"github.com/ulule/makroud"

	"git.iiens.net/edouardparis/town/failures"
	"git.iiens.net/edouardparis/town/models"
)

type Orders struct {
	Store
}

func (o Orders) GetByPublicID(ctx context.Context, publicID string) (*models.Order, error) {
	order := &models.Order{}

	q := lk.Select(columns(order)).
		From(order.TableName()).
		Where(lk.Condition("public_id").Equal(publicID))

	err := o.Get(ctx, q, order)
	if err != nil {
		if !makroud.IsErrNoRows(err) {
			return nil, errors.Wrapf(err, "cannot retrieve order with publicID: %s", publicID)
		}
		return nil, failures.ErrNotFound
	}

	return order, nil
}

func (o *Orders) Create(ctx context.Context, order *models.Order) error {
	query := lk.Insert(order.TableName()).
		Set(
			lk.Pair("public_id", order.PublicID),
			lk.Pair("description", order.Description),
			lk.Pair("amount", order.Amount),
			lk.Pair("status", order.Status),
			lk.Pair("fee", order.Fee),
			lk.Pair("fiat_value", order.FiatValue),
			lk.Pair("currency", order.Currency),
			lk.Pair("notes", order.Notes),
			lk.Pair("payreq", order.PayReq),
			lk.Pair("charge_created_at", order.ChargeCreatedAt),
			lk.Pair("charge_settled_at", order.ChargeSettledAt),
			lk.Pair("claimed_at", order.ClaimedAt),
		).Returning("id", "created_at", "updated_at")
	err := o.Save(ctx, query, order)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func NewOrders(s Store) *Orders {
	return &Orders{s}
}
