package front

import (
	"net/http"

	"github.com/go-chi/chi"
	chirender "github.com/go-chi/render"
	"github.com/pkg/errors"

	"git.iiens.net/edouardparis/town/failures"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/hi", func(w http.ResponseWriter, r *http.Request) {
		render(w, r, "<html><body><hi>Hi lizard</hi></body></html>", http.StatusOK)
	})
	return r
}

func render(w http.ResponseWriter, r *http.Request, resource string, httpStatus int) {
	if resource == "" && httpStatus == http.StatusNoContent {
		chirender.NoContent(w, r)
	}

	chirender.Status(r, httpStatus)
	chirender.HTML(w, r, resource)
}

type handlerFunc func(w http.ResponseWriter, r *http.Request) error

func handle(handler handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err == nil {
			return
		}

		var status int
		switch cerr := errors.Cause(err).(type) {
		case failures.Error:
			status = cerr.Code
			err = cerr
		default:
			status = http.StatusInternalServerError
		}

		chirender.Status(r, status)
		chirender.JSON(w, r, err)
	}
}
