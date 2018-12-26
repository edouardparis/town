package store

import (
	"context"

	"github.com/pkg/errors"
	lk "github.com/ulule/loukoum"
	"github.com/ulule/makroud"

	"git.iiens.net/edouardparis/town/failures"
	"git.iiens.net/edouardparis/town/models"
)

type Articles struct {
	store Store
}

func (a Articles) GetByID(ctx context.Context, id int64) (*models.Article, error) {
	article := &models.Article{}

	q := lk.Select(columns(article)).
		From(article.TableName()).
		Where(lk.Condition("id").Equal(id))

	err := a.store.Get(ctx, q, article)
	if err != nil {
		if !makroud.IsErrNoRows(err) {
			return nil, errors.Wrapf(err, "cannot retrieve article with ID: %d", id)
		}
		return nil, failures.ErrNotFound
	}

	err = a.Preload(ctx, article)
	if err != nil {
		return nil, err
	}

	return article, nil
}

// Preload populates the given article.
func (a *Articles) Preload(ctx context.Context, article *models.Article) error {
	if article == nil {
		return nil
	}

	articles := []models.Article{*article}

	err := a.PreloadList(ctx, &articles)
	if err != nil {
		return err
	}

	*article = articles[0]

	return nil
}

// PreloadList populates the given articles.
func (a *Articles) PreloadList(ctx context.Context, articles *[]models.Article) error {
	if articles == nil || len(*articles) == 0 {
		return nil
	}

	return makroud.Preload(ctx, a.store.Conn(), articles,
		makroud.WithPreloadField("Address"),
		makroud.WithPreloadField("Node"),
	)
}
