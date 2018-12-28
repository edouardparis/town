package managers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/avelino/slugify"
	"github.com/microcosm-cc/bluemonday"
	funk "github.com/thoas/go-funk"

	"git.iiens.net/edouardparis/town/constants"
	"git.iiens.net/edouardparis/town/failures"
	"git.iiens.net/edouardparis/town/models"
	"git.iiens.net/edouardparis/town/payloads"
	"git.iiens.net/edouardparis/town/store"
)

func ArticleCreate(ctx context.Context, s store.Store, payload *payloads.Article) (*models.Article, error) {
	sanitizeArticlePayload(payload)

	article := &models.Article{
		Title:  payload.Title,
		Body:   payload.Body,
		Lang:   constants.LangEN,
		Status: constants.ArticleStatusPublished,
	}

	if payload.Lang != "" {
		article.Lang = constants.LangStrToInt[payload.Lang]
	}

	err := s.Transaction(ctx, func(tx store.Store) error {
		err := setArticleSlug(ctx, tx, article, payload.Title)
		if err != nil {
			return err
		}

		if payload.Address != "" {
			err = setArticleAddress(ctx, tx, article, payload.Address)
		} else if payload.NodePubKey != "" {
			err = setArticleNode(ctx, tx, article, payload.NodePubKey)
		}
		if err != nil {
			return err
		}

		return store.NewArticles(tx).Create(ctx, article)
	})
	if err != nil {
		return nil, err
	}

	return article, nil
}

func setArticleSlug(ctx context.Context, s store.Store, article *models.Article, title string) error {
	candidat := slugify.Slugify(title)
	slugStore := store.NewSlugs(s)
	for {
		s, err := slugStore.GetBySlug(ctx, candidat)
		if err == failures.ErrNotFound && s == nil {
			break
		} else if err != nil {
			return err
		}
		candidat = fmt.Sprintf("%s-%s", candidat, funk.RandomString(5))
	}
	slug := &models.Slug{Slug: candidat}
	err := slugStore.Create(ctx, slug)
	if err != nil {
		return err
	}

	article.Slug = slug.Slug
	return nil
}

func setArticleAddress(ctx context.Context, s store.Store, article *models.Article, value string) error {
	addressStore := store.NewAddresses(s)
	address, err := addressStore.GetByValue(ctx, value)
	if err != nil || address == nil {
		if err != failures.ErrNotFound {
			return err
		} else {
			address = &models.Address{
				Value: value,
			}
			err := addressStore.Create(ctx, address)
			if err != nil {
				return err
			}
		}
	}

	article.Address = address
	article.AddressID = sql.NullInt64{
		Int64: address.ID,
		Valid: true,
	}

	return nil
}

func setArticleNode(ctx context.Context, s store.Store, article *models.Article, pubKey string) error {
	node := &models.Node{}
	err := store.NewNodes(s).Create(ctx, node)
	if err != nil {
		return err
	}
	article.Node = node
	article.NodeID = sql.NullInt64{
		Int64: node.ID,
		Valid: true,
	}
	return nil
}

func sanitizeArticlePayload(payload *payloads.Article) {
	payload.Title = bluemonday.UGCPolicy().Sanitize(payload.Title)
	payload.Subtitle = bluemonday.UGCPolicy().Sanitize(payload.Subtitle)
	payload.Body = bluemonday.UGCPolicy().Sanitize(payload.Body)
}
