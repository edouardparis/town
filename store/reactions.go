package store

import (
	"context"

	"github.com/pkg/errors"
	lk "github.com/ulule/loukoum"

	"github.com/EdouardParis/town/models"
)

type Reactions struct {
	Store
}

func NewReactions(s Store) *Reactions {
	return &Reactions{s}
}

func (a *Reactions) Create(ctx context.Context, reaction *models.Reaction) error {
	query := lk.Insert(reaction.TableName()).
		Set(
			lk.Pair("article_id", reaction.ArticleID),
			lk.Pair("order_id", reaction.OrderID),
			lk.Pair("emoji", reaction.Emoji),
		).Returning("id", "created_at")
	err := a.Save(ctx, query, reaction)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
