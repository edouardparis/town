package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/EdouardParis/town/app"
)

// Run the server
func Run(ctx context.Context, app *app.App) error {
	g, ctx := errgroup.WithContext(ctx)

	// Signal handler
	g.Go(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		select {
		case s := <-c:
			app.Logger.Info("Interrupt peacefully server")
			return fmt.Errorf("Got signal %v", s)
		case <-ctx.Done():
			return nil
		}
	})

	// HTTP server
	g.Go(func() error {
		cerr := make(chan error)
		srv := http.Server{
			Addr:    ":8080",
			Handler: Routes(ctx, app),
		}

		app.Logger.Info("listening on port :8080")
		go func() { cerr <- errors.WithStack(srv.ListenAndServe()) }()
		select {
		case err := <-cerr:
			return err
		case <-ctx.Done():
			return srv.Shutdown(context.Background())
		}
	})

	return g.Wait()
}
