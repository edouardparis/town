package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"

	"git.iiens.net/edouardparis/town/app"
	"git.iiens.net/edouardparis/town/web/api"
	"git.iiens.net/edouardparis/town/web/front"
)

// Routes register all the API urls and handlers to the router.
func Routes(ctx context.Context, a *app.App) http.Handler {
	r := chi.NewRouter()
	r.Mount("/", front.NewRouter(a))
	r.Mount("/api", api.NewRouter(a))
	return chi.ServerBaseContext(ctx, r)
}
