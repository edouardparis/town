package front

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	chirender "github.com/go-chi/render"
	"github.com/pkg/errors"

	"git.iiens.net/edouardparis/town/app"
	"git.iiens.net/edouardparis/town/failures"
	"git.iiens.net/edouardparis/town/templates"
)

func NewRouter(a *app.App) http.Handler {
	r := chi.NewRouter()
	r.Get("/", handle(home))
	r.Route("/articles", articlesRoutes(a))

	workDir, _ := os.Getwd()
	FileServer(r, "/static", http.Dir(filepath.Join(workDir, "static")))

	return r
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		return
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
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
