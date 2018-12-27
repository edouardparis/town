package front

import (
	"net/http"

	"git.iiens.net/edouardparis/town/app"
	"git.iiens.net/edouardparis/town/failures"
	"git.iiens.net/edouardparis/town/resources"
	"git.iiens.net/edouardparis/town/web/middlewares"
)

func Home(a *app.App, handle middlewares.HandleError) http.HandlerFunc {
	data := struct {
		Header *resources.Header
	}{}
	return func(w http.ResponseWriter, r *http.Request) {
		data.Header = resources.NewHeader(a.Info)
		err := render(w, r, "home.html", data, http.StatusOK)
		if err != nil {
			handle(w, r, failures.ErrNotFound)
		}
	}
}
