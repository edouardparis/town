package managers

import (
	"context"
	"time"

	"github.com/microcosm-cc/bluemonday"

	"git.iiens.net/edouardparis/town/constants"
	"git.iiens.net/edouardparis/town/models"
	"git.iiens.net/edouardparis/town/payloads"
	"git.iiens.net/edouardparis/town/store"
)

func ArticleCreate(ctx context.Context, s store.Store, payload *payloads.Article) (*models.Article, error) {
	sanitizeArticlePayload(payload)

	article := &models.Article{
		Title:     payload.Title,
		Body:      payload.Body,
		Lang:      constants.LangEN,
		Status:    constants.ArticleStatusPublished,
		CreatedAt: time.Now().UTC(),
	}

	if payload.Lang != "" {
		article.Lang = constants.LangStrToInt[payload.Lang]
	}

	err := s.Transaction(ctx, func(tx store.Store) error {
		return store.NewArticles(tx).Create(ctx, article)
	})
	if err != nil {
		return nil, err
	}

	return article, nil
}

func sanitizeArticlePayload(payload *payloads.Article) {
	payload.Title = bluemonday.UGCPolicy().Sanitize(payload.Title)
	payload.Subtitle = bluemonday.UGCPolicy().Sanitize(payload.Subtitle)
	payload.Body = bluemonday.UGCPolicy().Sanitize(payload.Body)
}
