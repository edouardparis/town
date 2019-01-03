package models

import (
	"time"

	"github.com/lib/pq"
)

type Order struct {
	ID       string `makroud:"column:id"`
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
	FiatValue int64 `makroud:"column:fiat_value"`
	// Charge currency
	Currency string `makroud:"column:currency"`
	// Charge notes
	Notes string `makroud:"column:notes"`

	// Payment Request hash
	PayReq string `makroud:"column:payreq"`

	ChargeCreatedAt time.Time `makroud:"column:charge_created_at"`
	ChargeSettledAt time.Time `makroud:"column:charge_settle_at"`

	CreatedAt time.Time   `makroud:"column:created_at"`
	UpdatedAt pq.NullTime `makroud:"column:updated_at"`
	ClaimedAt pq.NullTime `makroud:"column:claimed_at"`
}

// TableName implements Model interface.
func (Order) TableName() string {
	return "town_order"
}
