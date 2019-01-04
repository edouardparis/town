package store

import (
	"context"

	"github.com/pkg/errors"
	lk "github.com/ulule/loukoum"
	"github.com/ulule/makroud"

	"github.com/EdouardParis/town/constants"
	"github.com/EdouardParis/town/failures"
	"github.com/EdouardParis/town/models"
)

type Articles struct {
	Store
}

func NewArticles(s Store) *Articles {
	return &Articles{s}
}

func (a Articles) GetByID(ctx context.Context, id int64) (*models.Article, error) {
	article := &models.Article{}

	q := lk.Select(columns(article)).
		From(article.TableName()).
		Where(lk.Condition("id").Equal(id))

	err := a.Get(ctx, q, article)
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

func (a Articles) GetBySlug(ctx context.Context, slug string) (*models.Article, error) {
	article := &models.Article{}

	q := lk.Select(columns(article)).
		From(article.TableName()).
		Where(lk.Condition("slug").Equal(slug))

	err := a.Get(ctx, q, article)
	if err != nil {
		if !makroud.IsErrNoRows(err) {
			return nil, errors.Wrapf(err, "cannot retrieve article with ID: %s", slug)
		}
		return nil, failures.ErrNotFound
	}

	err = a.Preload(ctx, article)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (a Articles) ListPublished(ctx context.Context) ([]models.Article, error) {
	articles := []models.Article{}

	model := models.Article{}
	q := lk.Select(columns(model)).
		From(model.TableName()).
		Where(lk.Condition("status").Equal(constants.ArticleStatusPublished))

	err := a.Find(ctx, q, &articles)
	if err != nil {
		if !makroud.IsErrNoRows(err) {
			return nil, errors.Wrapf(err, "cannot retrieve published articles")
		}
		return nil, failures.ErrNotFound
	}

	err = a.PreloadList(ctx, &articles)
	if err != nil {
		return nil, err
	}

	return articles, nil
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

	return makroud.Preload(ctx, a.Conn(), articles,
		makroud.WithPreloadField("Address"),
		makroud.WithPreloadField("Node"),
	)
}

func (a *Articles) Create(ctx context.Context, article *models.Article) error {
	query := lk.Insert(article.TableName()).
		Set(
			lk.Pair("title", article.Title),
			lk.Pair("subtitle", article.Subtitle),
			lk.Pair("body", article.Body),
			lk.Pair("lang", article.Lang),
			lk.Pair("amount_collected", article.AmountCollected),
			lk.Pair("amount_received", article.AmountReceived),
			lk.Pair("status", article.Status),
			lk.Pair("slug", article.Slug),
			lk.Pair("published_at", article.PublishedAt),
			lk.Pair("address_id", article.AddressID),
			lk.Pair("node_id", article.NodeID),
		).Returning("id", "created_at", "updated_at")
	err := a.Save(ctx, query, article)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
