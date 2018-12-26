package models

import (
	"time"

	"github.com/lib/pq"
)

type Node struct {
	ID              int64       `makroud:"column:id,pk"`
	PubKey          string      `makroud:"column:pub_key"`
	AmountCollected int64       `makroud:"column:amount_collected"`
	AmountReceived  int64       `makroud:"column:amount_received"`
	CreatedAt       time.Time   `makroud:"column:created_at"`
	UpdatedAt       pq.NullTime `makroud:"column:updated_at"`
}

// TableName implements Model interface.
func (Node) TableName() string {
	return "town_node"
}
