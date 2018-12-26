package front

import (
	"net/http"

	"github.com/go-chi/chi"
	chirender "github.com/go-chi/render"
	"github.com/pkg/errors"

	"git.iiens.net/edouardparis/town/failures"
	"git.iiens.net/edouardparis/town/templates"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", handle(home))
	return r
}

func render(w http.ResponseWriter, r *http.Request, template string, resource interface{}, httpStatus int) error {
	err := templates.HTMLTemplates.ExecuteTemplate(w, template, resource)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	chirender.Status(r, httpStatus)
	return nil
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
