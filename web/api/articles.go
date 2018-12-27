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
	handle := newHandle(a)
	return func(r chi.Router) {
		r.Route("/{slug:[a-z-]+}", func(r chi.Router) {
			r.With(middlewares.ArticleCtx(a, handleError)).
				Get("/", handle(ArticleDetail))

			r.Post("/", handle(ArticleCreate))
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

func ArticleCreate(a *app.App, handle middlewares.HandleError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := &payloads.Article{}
		errs := binding.Bind(r, payload)
		if errs != nil {
			handle(w, r, errs)
			return
		}

		article, err := managers.ArticleCreate(r.Context(), a.Store, payload)
		if err != nil {
			handle(w, r, err)
			return
		}

		resource := resources.NewArticle(article)
		err = render(w, r, resource, http.StatusCreated)
		if err != nil {
			handle(w, r, err)
		}
	}
}
