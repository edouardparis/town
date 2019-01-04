package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	chirender "github.com/go-chi/render"
	"github.com/mholt/binding"
	"github.com/pkg/errors"

	"git.iiens.net/edouardparis/town/app"
	"git.iiens.net/edouardparis/town/failures"
	"git.iiens.net/edouardparis/town/logging"
)

func NewRouter(a *app.App) http.Handler {
	r := chi.NewRouter()
	r.Route("/articles", articlesRoutes(a))
	r.Route("/webhooks", webhooksRoutes(a))
	return r
}

func render(w http.ResponseWriter, r *http.Request, resource interface{}, httpStatus int) error {
	chirender.Status(r, httpStatus)

	if resource == nil {
		return nil
	}

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	err := enc.Encode(resource)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = w.Write(buf.Bytes())
	return err
}

func handleError(logger logging.Logger) func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		if err == nil {
			return
		}

		var status int
		switch cerr := errors.Cause(err).(type) {
		case failures.Error:
			status = cerr.Code
			err = cerr
		case binding.Errors:
			status = http.StatusBadRequest
			err = cerr
		default:
			logger.Error(cerr.Error())
			status = http.StatusInternalServerError
		}

		chirender.Status(r, status)
		chirender.JSON(w, r, err)
	}
}

func handle(a *app.App, fn func(*app.App) func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(a)(w, r)
		if err == nil {
			return
		}

		var status int
		switch cerr := errors.Cause(err).(type) {
		case failures.Error:
			status = cerr.Code
			err = cerr
		case binding.Errors:
			status = http.StatusBadRequest
			err = cerr
		default:
			a.Logger.Error(cerr.Error())
			status = http.StatusInternalServerError
		}

		chirender.Status(r, status)
		chirender.JSON(w, r, err)
	}
}
