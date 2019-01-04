package front

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/EdouardParis/town/app"
	"github.com/EdouardParis/town/failures"
	"github.com/EdouardParis/town/resources"
	"github.com/EdouardParis/town/store"
	"github.com/EdouardParis/town/web/middlewares"
)

func articlesRoutes(a *app.App) func(r chi.Router) {
	handle := newHandle(a)
	articleCtx := middlewares.ArticleCtx(a, handleError(a.Logger))
	return func(r chi.Router) {
		r.Get("/", handle(ArticleList))
		r.Get("/write", handle(ArticleWrite))
		r.Route("/{slug}", func(r chi.Router) {
			r.With(articleCtx).Get("/", handle(ArticleDetail))
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
