package api

import (
	"net/http"

	"github.com/go-chi/chi"

	"git.iiens.net/edouardparis/town/app"
	"git.iiens.net/edouardparis/town/failures"
	"git.iiens.net/edouardparis/town/resources"
	"git.iiens.net/edouardparis/town/web/middlewares"
)

func articlesRoutes(a *app.App) func(r chi.Router) {
	handle := newHandle(a)
	return func(r chi.Router) {
		r.Route("/{slug:[a-z-]+}", func(r chi.Router) {
			r.Use(middlewares.ArticleCtx(a, handleError))
			r.Get("/", handle(ArticleDetail))
		})
	}
}

func ArticleDetail(a *app.App, handle middlewares.HandleError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		article, ok := middlewares.ArticleFromCtx(ctx)
		if !ok {
			handle(w, r, failures.ErrNotFound)
			return
		}
		resource := resources.NewArticle(article)
		err := render(w, r, resource, http.StatusOK)
		if err != nil {
			handle(w, r, err)
		}
	}
}
