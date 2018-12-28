package middlewares

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"

	"git.iiens.net/edouardparis/town/app"
	"git.iiens.net/edouardparis/town/failures"
	"git.iiens.net/edouardparis/town/models"
	"git.iiens.net/edouardparis/town/store"
)

type contextKey string

func (c contextKey) String() string {
	return "middlewares " + string(c)
}

var (
	articleKey = contextKey("article")
	addressKey = contextKey("address")
)

func ArticleFromCtx(ctx context.Context) (*models.Article, bool) {
	article, ok := ctx.Value(articleKey).(*models.Article)
	return article, ok
}

func ArticleCtx(a *app.App, handle HandleError) Middleware {
	s := store.NewArticles(a.Store)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			slug := chi.URLParam(r, "slug")
			article, err := s.GetBySlug(r.Context(), slug)
			if err != nil {
				handle(w, r, err)
				return
			}

			if !article.IsPublished() {
				handle(w, r, failures.ErrNotFound)
				return
			}
			ctx := context.WithValue(r.Context(), articleKey, article)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AddressFromCtx(ctx context.Context) (*models.Address, bool) {
	address, ok := ctx.Value(addressKey).(*models.Address)
	return address, ok
}

func AddressCtx(a *app.App, handle HandleError) Middleware {
	s := store.NewAddresses(a.Store)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			slug := chi.URLParam(r, "value")
			address, err := s.GetByValue(r.Context(), slug)
			if err != nil {
				handle(w, r, err)
				return
			}

			ctx := context.WithValue(r.Context(), addressKey, address)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
