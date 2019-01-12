package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mholt/binding"

	"github.com/EdouardParis/town/app"
	"github.com/EdouardParis/town/failures"
	"github.com/EdouardParis/town/managers"
	"github.com/EdouardParis/town/payloads"
	"github.com/EdouardParis/town/resources"
	"github.com/EdouardParis/town/web/middlewares"
)

func articlesRoutes(a *app.App) func(r chi.Router) {
	articleCtx := middlewares.ArticleCtx(a, handleError(a.Logger))
	return func(r chi.Router) {
		r.Route("/{slug:[a-z-]+}", func(r chi.Router) {
			r.Use(articleCtx)
			r.Get("/", handle(a, ArticleDetail))

			r.Route("/reactions", reactionsRoutes(a))
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
