package store

import (
	"context"

	"github.com/pkg/errors"
	lk "github.com/ulule/loukoum"
	"github.com/ulule/makroud"

	"github.com/EdouardParis/town/failures"
	"github.com/EdouardParis/town/models"
)

type Slugs struct {
	Store
}

func NewSlugs(s Store) *Slugs {
	return &Slugs{s}
}

func (a Slugs) GetBySlug(ctx context.Context, s string) (*models.Slug, error) {
	slug := &models.Slug{}

	q := lk.Select(columns(slug)).
		From(slug.TableName()).
		Where(lk.Condition("slug").Equal(s))

	err := a.Get(ctx, q, slug)
	if err != nil {
		if !makroud.IsErrNoRows(err) {
			return nil, errors.Wrapf(err, "cannot retrieve slug with value: %s", s)
		}
		return nil, failures.ErrNotFound
	}

	return slug, nil
}

func (a *Slugs) Create(ctx context.Context, slug *models.Slug) error {
	query := lk.Insert(slug.TableName()).
		Set(
			lk.Pair("slug", slug.Slug),
			lk.Pair("current_id", slug.CurrentID),
		).Returning("id", "created_at", "updated_at")
	err := a.Save(ctx, query, slug)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
