package managers

import (
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"

	"git.iiens.net/edouardparis/town/opennode"
	"git.iiens.net/edouardparis/town/resources"
)

var _ = uuid.Must(uuid.NewV4())

func CreateCharge(c *opennode.Config) (*resources.Charge, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, errors.Wrap(err, "Error during uuid creation")
	}

	charge, err := opennode.NewClient(c).CreateCharge(&opennode.ChargePayload{
		Amount:   int64(6),
		Currency: "EUR",
		OrderID:  uid.String(),
	})

	if err != nil {
		return nil, errors.Wrap(err, "Error during charge creation")
	}

	return resources.NewCharge(charge), nil
}
