package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/EdouardParis/town/app"
	"github.com/EdouardParis/town/web/api"
	"github.com/EdouardParis/town/web/front"
	"github.com/EdouardParis/town/web/middlewares"
	"github.com/EdouardParis/town/web/websockets"
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
