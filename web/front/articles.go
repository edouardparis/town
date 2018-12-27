package front

import (
	"net/http"

	"github.com/go-chi/chi"

	"git.iiens.net/edouardparis/town/app"
	"git.iiens.net/edouardparis/town/failures"
	"git.iiens.net/edouardparis/town/resources"
	"git.iiens.net/edouardparis/town/store"
	"git.iiens.net/edouardparis/town/web/middlewares"
)

func articlesRoutes(a *app.App) func(r chi.Router) {
	handle := newHandle(a)
	return func(r chi.Router) {
		r.Get("/", handle(ArticleList))
		r.Get("/write", handle(ArticleWrite))
		r.Route("/{slug:[a-z-]+}", func(r chi.Router) {
			r.Use(middlewares.ArticleCtx(a, handleError))
			r.Get("/", handle(ArticleDetail))
		})
	}
}

func ArticleDetail(a *app.App, handle middlewares.HandleError) http.HandlerFunc {
	data := struct {
		Header  *resources.Header
		Article *resources.Article
	}{}
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		article, ok := middlewares.ArticleFromCtx(ctx)
		if !ok || article == nil {
			handle(w, r, failures.ErrNotFound)
			return
		}
		data.Article = resources.NewArticle(article)
		data.Header = resources.NewHeader(a.Info)
		err := render(w, r, "article.html", data, http.StatusOK)
		if err != nil {
			handle(w, r, failures.ErrNotFound)
		}
	}
}

func ArticleList(a *app.App, handle middlewares.HandleError) http.HandlerFunc {
	data := struct {
		Articles []resources.Article
		Header   *resources.Header
	}{Header: resources.NewHeader(a.Info)}
	s := store.NewArticles(a.Store)
	return func(w http.ResponseWriter, r *http.Request) {
		articles, err := s.ListPublished(r.Context())
		if err != nil {
			handle(w, r, err)
			return
		}

		data.Articles = resources.NewArticleList(articles)

		err = render(w, r, "articles.html", data, http.StatusOK)
		if err != nil {
			handle(w, r, failures.ErrNotFound)
		}
	}
}

func ArticleWrite(a *app.App, handle middlewares.HandleError) http.HandlerFunc {
	data := struct{ Header *resources.Header }{Header: resources.NewHeader(a.Info)}
	return func(w http.ResponseWriter, r *http.Request) {
		err := render(w, r, "write.html", data, http.StatusOK)
		if err != nil {
			handle(w, r, failures.ErrNotFound)
		}
	}
}
