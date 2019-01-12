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

func reactionsRoutes(a *app.App) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", handle(a, ReactionCreate))
	}
}

func ReactionCreate(a *app.App) func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		payload := &payloads.Reaction{}
		errs := binding.Bind(r, payload)
		if errs != nil {
			return errs
		}

		ctx := r.Context()
		article, ok := middlewares.ArticleFromCtx(ctx)
		if !ok {
			return failures.ErrNotFound
		}

		reaction, err := managers.ReactionCreate(ctx, a.Store, payload, article)
		if err != nil {
			return err
		}

		resource := resources.NewReaction(reaction)
		return render(w, r, resource, http.StatusCreated)
	}
}
