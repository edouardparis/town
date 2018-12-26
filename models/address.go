package models

import (
	"time"

	"github.com/lib/pq"
)

type Address struct {
	ID              int64       `makroud:"column:id,pk"`
	Value           string      `makroud:"column:value"`
	AmountCollected int64       `makroud:"column:amount_collected"`
	AmountReceived  int64       `makroud:"column:amount_received"`
	CreatedAt       time.Time   `makroud:"column:created_at"`
	UpdatedAt       pq.NullTime `makroud:"column:updated_at"`
}

// TableName implements Model interface.
func (Address) TableName() string {
	return "town_address"
}
