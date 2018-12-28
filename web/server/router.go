package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"git.iiens.net/edouardparis/town/app"
	"git.iiens.net/edouardparis/town/web/api"
	"git.iiens.net/edouardparis/town/web/front"
	"git.iiens.net/edouardparis/town/web/middlewares"
	"git.iiens.net/edouardparis/town/web/websockets"
)

// Routes register all the API urls and handlers to the router.
func Routes(ctx context.Context, a *app.App) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middlewares.Logger(a.Logger))
	r.Mount("/", front.NewRouter(a))
	r.Mount("/api", api.NewRouter(a))
	r.Mount("/ws", websockets.NewRouter(a))
	return chi.ServerBaseContext(ctx, r)
}
