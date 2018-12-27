package front

import (
	"net/http"

	"git.iiens.net/edouardparis/town/app"
	"git.iiens.net/edouardparis/town/failures"
	"git.iiens.net/edouardparis/town/web/middlewares"
)

func Home(a *app.App, handle middlewares.HandleError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		home := struct {
			Title string
		}{
			Title: "HEllo",
		}
		err := render(w, r, "home.html", home, http.StatusOK)
		if err != nil {
			handle(w, r, failures.ErrNotFound)
		}
	}
}
