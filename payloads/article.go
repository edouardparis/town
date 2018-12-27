package payloads

import (
	"net/http"

	"github.com/mholt/binding"
)

type Article struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Body     string `json:"body"`
	Lang     string `json:"lang"`
}

// FieldMap for payload (github.com/mholt/binding)
func (a *Article) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&a.Title:    "title",
		&a.Subtitle: "subtitle",
		&a.Body:     "body",
		&a.Lang:     "lang",
	}
}

func (a *Article) Validate(req *http.Request) error {
	return nil
}
