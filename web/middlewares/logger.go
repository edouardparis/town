package middlewares

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"

	"github.com/EdouardParis/town/logging"
)

func Logger(logger logging.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fields := []logging.Field{
				logging.String("method", r.Method),
				logging.String("path", r.URL.Path),
				logging.String("uri", r.RequestURI),
			}

			if addr, port, err := net.SplitHostPort(r.RemoteAddr); err == nil {
				fields = append(fields, logging.String("remote-addr", addr), logging.String("remote-port", port))
			}

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			start := time.Now()

			next.ServeHTTP(ww, r)

			fields = append(fields,
				logging.Duration("duration", time.Since(start)),
				logging.Int("status", ww.Status()),
			)

			logger.Info(fmt.Sprintf("%d %s %s", ww.Status(), r.Method, r.URL.Path), fields...)
		})
	}
}
