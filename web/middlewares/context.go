package middlewares

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"

	"git.iiens.net/edouardparis/town/app"
	"git.iiens.net/edouardparis/town/store"
)

func ArticleCtx(a *app.App) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			slug := chi.URLParam(r, "slug")
			article, err := store.NewArticles(a.Store).GetBySlug(r.Context(), slug)
			if err != nil {
				http.Error(w, http.StatusText(404), 404)
				return
			}
			ctx := context.WithValue(r.Context(), "article", article)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
