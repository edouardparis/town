package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

// Routes register all the API urls and handlers to the router.
func Routes(ctx context.Context) http.Handler {
	r := chi.NewRouter()
	return chi.ServerBaseContext(ctx, r)
}
