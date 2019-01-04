package api

import (
	"net/http"

	"github.com/go-chi/chi"

	"git.iiens.net/edouardparis/town/app"
	"git.iiens.net/edouardparis/town/failures"
	"git.iiens.net/edouardparis/town/resources"
	"git.iiens.net/edouardparis/town/web/middlewares"
)

func webhooksRoutes(a *app.App) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/checkout", handle(a, CheckoutWebhook))
	}
}

func CheckoutWebhook(a *app.App) func(http.ResponseWriter, *http.Request) error {
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
