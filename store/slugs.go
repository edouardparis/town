package store

import (
	"context"

	"github.com/pkg/errors"
	lk "github.com/ulule/loukoum"

	"git.iiens.net/edouardparis/town/models"
)

type Slugs struct {
	Store
}

func NewSlugs(s Store) *Slugs {
	return &Slugs{s}
}

func (a *Slugs) Create(ctx context.Context, slug *models.Slug) error {
	query := lk.Insert(slug.TableName()).
		Set(
			lk.Pair("slug", slug.Slug),
			lk.Pair("current_id", slug.CurrentID),
			lk.Pair("created_at", slug.CreatedAt),
			lk.Pair("updated_at", slug.UpdatedAt),
		).Returning("id")
	err := a.Save(ctx, query, slug)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
