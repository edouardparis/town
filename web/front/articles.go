package front

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"git.iiens.net/edouardparis/town/app"
	"git.iiens.net/edouardparis/town/models"
	"git.iiens.net/edouardparis/town/web/middlewares"
)

func articlesRoutes(a *app.App) func(r chi.Router) {
	return func(r chi.Router) {
		r.Route("/{slug:[a-z-]+}", func(r chi.Router) {
			r.Use(middlewares.ArticleCtx(a))
			r.Get("/", ArticleDetail(a))
		})
	}
}

func ArticleDetail(a *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		article, ok := ctx.Value("article").(*models.Article)
		if !ok {
			http.Error(w, http.StatusText(422), 422)
			return
		}
		w.Write([]byte(fmt.Sprintf("title:%s", article.Title)))
	}
}
