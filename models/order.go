package models

import (
	"time"

	"github.com/EdouardParis/town/constants"
	"github.com/LizardsTown/opennode"
	"github.com/lib/pq"
)

type Order struct {
	ID       int64  `makroud:"column:id,pk"`
	PublicID string `makroud:"column:public_id"`
	// Charge Description
	Description string `makroud:"column:description"`
	// Charge price in satoshis
	Amount int64 `makroud:"column:amount"`
	// Charge status
	// unpaid/processing/paid/claimed
	Status int64 `makroud:"column:status"`
	// Charge fee in satoshis
	Fee int64 `makroud:"column:fee"`
	// Charge value at issue time
	FiatValue float64 `makroud:"column:fiat_value"`
	// Charge currency
	Currency int64 `makroud:"column:currency"`
	// Charge notes
	Notes string `makroud:"column:notes"`

	// Payment Request hash
	PayReq string `makroud:"column:payreq"`

	ChargeCreatedAt time.Time `makroud:"column:charge_created_at"`
	ChargeSettleAt  time.Time `makroud:"column:charge_settle_at"`

	CreatedAt time.Time   `makroud:"column:created_at"`
	UpdatedAt pq.NullTime `makroud:"column:updated_at"`
	ClaimedAt pq.NullTime `makroud:"column:claimed_at"`
}

// TableName implements Model interface.
func (Order) TableName() string {
	return "town_order"
}

func NewOrder(c *opennode.Charge) *Order {
	return &Order{
		PublicID:        c.OrderID,
		Description:     c.Description,
		Amount:          c.Amount,
		Status:          constants.OrderStatusStrToInt[c.Status],
		Fee:             c.Fee,
		FiatValue:       c.FiatValue,
		Currency:        constants.CurrenciesStrToInt[c.Currency],
		Notes:           c.Notes,
		PayReq:          c.LightningInvoice.PayReq,
		ChargeCreatedAt: time.Unix(c.LightningInvoice.CreatedAt, 0),
		ChargeSettleAt:  time.Unix(c.LightningInvoice.SettledAt, 0),
	}
}
