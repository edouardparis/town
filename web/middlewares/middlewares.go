package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler
type HandleError func(http.ResponseWriter, *http.Request, error)
