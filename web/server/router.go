package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"

	"git.iiens.net/edouardparis/town/app"
)

// Routes register all the API urls and handlers to the router.
func Routes(ctx context.Context, app *app.App) http.Handler {
	r := chi.NewRouter()
	return chi.ServerBaseContext(ctx, r)
}
