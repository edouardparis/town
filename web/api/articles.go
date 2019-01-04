package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mholt/binding"

	"git.iiens.net/edouardparis/town/app"
	"git.iiens.net/edouardparis/town/failures"
	"git.iiens.net/edouardparis/town/managers"
	"git.iiens.net/edouardparis/town/payloads"
	"git.iiens.net/edouardparis/town/resources"
	"git.iiens.net/edouardparis/town/web/middlewares"
)

func articlesRoutes(a *app.App) func(r chi.Router) {
	articleCtx := middlewares.ArticleCtx(a, handleError(a.Logger))
	return func(r chi.Router) {
		r.Route("/{slug:[a-z-]+}", func(r chi.Router) {
			r.With(articleCtx).Get("/", handle(a, ArticleDetail))
		})

		r.Post("/", handle(a, ArticleCreate))
	}
}

func ArticleDetail(a *app.App) func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		article, ok := middlewares.ArticleFromCtx(ctx)
		if !ok {
			return failures.ErrNotFound
		}
		resource := resources.NewArticle(article)
		return render(w, r, resource, http.StatusOK)
	}
}

func ArticleCreate(a *app.App) func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		payload := &payloads.Article{}
		errs := binding.Bind(r, payload)
		if errs != nil {
			return errs
		}

		article, err := managers.ArticleCreate(r.Context(), a.Store, payload)
		if err != nil {
			return err
		}

		resource := resources.NewArticle(article)
		return render(w, r, resource, http.StatusCreated)
	}
}
