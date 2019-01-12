package store

import (
	"context"

	"github.com/pkg/errors"
	lk "github.com/ulule/loukoum"

	"github.com/EdouardParis/town/models"
)

type Comments struct {
	Store
}

func NewComments(s Store) *Comments {
	return &Comments{s}
}

func (a *Comments) Create(ctx context.Context, comment *models.ArticleComment) error {
	query := lk.Insert(comment.TableName()).
		Set(
			lk.Pair("article_id", comment.ArticleID),
			lk.Pair("order_id", comment.OrderID),
			lk.Pair("body", comment.Body),
		).Returning("id", "created_at")
	err := a.Save(ctx, query, comment)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
