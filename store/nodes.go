package store

import (
	"context"

	"github.com/pkg/errors"
	lk "github.com/ulule/loukoum"

	"git.iiens.net/edouardparis/town/models"
)

type Nodes struct {
	Store
}

func NewNodes(s Store) *Nodes {
	return &Nodes{s}
}

func (a *Nodes) Create(ctx context.Context, node *models.Node) error {
	query := lk.Insert(node.TableName()).
		Set(
			lk.Pair("pub_key", node.PubKey),
			lk.Pair("amount_collected", node.AmountCollected),
			lk.Pair("amount_received", node.AmountReceived),
		).Returning("id", "created_at", "updated_at")
	err := a.Save(ctx, query, node)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
