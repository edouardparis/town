package resources

import (
	"time"

	"git.iiens.net/edouardparis/town/models"
)

type Address struct {
	Value           string     `json:"value"`
	AmountCollected int64      `json:"amount_collected"`
	AmountReceived  int64      `json:"amount_received"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}

func NewAddress(address *models.Address) *Address {
	if address == nil {
		return nil
	}

	resource := &Address{
		Value:           address.Value,
		AmountCollected: address.AmountCollected,
		AmountReceived:  address.AmountReceived,
		CreatedAt:       address.CreatedAt,
	}

	if address.UpdatedAt.Valid {
		resource.UpdatedAt = &address.UpdatedAt.Time
	}

	return resource
}
