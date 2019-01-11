package payloads

import (
	"net/http"

	"github.com/mholt/binding"

	"github.com/EdouardParis/town/constants"
)

type Article struct {
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	Body       string `json:"body"`
	Lang       string `json:"lang"`
	Address    string `json:"address"`
	NodePubKey string `json:"node_pub_key"`
	OrderID    string `json:"order_id"`
}

// FieldMap for payload (github.com/mholt/binding)
func (a *Article) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&a.Title:      binding.Field{Form: "title", Required: true},
		&a.Subtitle:   "subtitle",
		&a.Body:       binding.Field{Form: "body", Required: true},
		&a.Lang:       "lang",
		&a.Address:    "address",
		&a.NodePubKey: "node_pub_key",
		&a.OrderID:    binding.Field{Form: "order_id", Required: true},
	}
}

func (a *Article) Validate(req *http.Request) error {
	if a.Lang != "" {
		if _, ok := constants.LangStrToInt[a.Lang]; !ok {
			return binding.Errors{
				binding.Error{
					FieldNames:     []string{"lang"},
					Classification: "BadValue",
					Message:        "lang must be en or fr",
				},
			}
		}
	}
	return nil
}
