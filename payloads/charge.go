package payloads

import (
	"net/http"

	"git.iiens.net/edouardparis/town/opennode"
	"github.com/mholt/binding"
)

type Charge struct {
	opennode.Charge
}

// FieldMap for payload (github.com/mholt/binding)
func (c *Charge) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&c.ID:          "id",
		&c.Status:      "status",
		&c.OrderID:     "order_id",
		&c.HashedOrder: "hashed_order",
	}
}
