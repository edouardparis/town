package front

import (
	"net/http"

	"github.com/EdouardParis/town/app"
	"github.com/EdouardParis/town/failures"
	"github.com/EdouardParis/town/resources"
	"github.com/EdouardParis/town/store"
	"github.com/EdouardParis/town/web/middlewares"
)

func Home(a *app.App, handle middlewares.HandleError) http.HandlerFunc {
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

		err = render(w, r, "home.html", data, http.StatusOK)
		if err != nil {
			handle(w, r, failures.ErrNotFound)
		}
	}
}

func About(a *app.App, handle middlewares.HandleError) http.HandlerFunc {
	data := struct {
		Header *resources.Header
	}{Header: resources.NewHeader(a.Info)}
	return func(w http.ResponseWriter, r *http.Request) {
		err := render(w, r, "about.html", data, http.StatusOK)
		if err != nil {
			handle(w, r, err)
		}
	}
}
