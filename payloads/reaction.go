package payloads

import (
	"net/http"

	"github.com/hackebrot/turtle"
	"github.com/mholt/binding"
)

type Reaction struct {
	Emoji   string `json:"emoji"`
	OrderID string `json:"order_id"`
}

// FieldMap for payload (github.com/mholt/binding)
func (a *Reaction) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&a.Emoji:   binding.Field{Form: "emoji", Required: true},
		&a.OrderID: binding.Field{Form: "order_id", Required: true},
	}
}

func (a *Reaction) Validate(req *http.Request) error {
	if a.Emoji != "" {
		if _, ok := turtle.Emojis[a.Emoji]; !ok {
			return binding.Errors{
				binding.Error{
					FieldNames:     []string{"emoji"},
					Classification: "BadValue",
					Message:        "emoji does not exist",
				},
			}
		}
	}
	return nil
}
