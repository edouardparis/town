package models

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Slug struct {
	ID        int64         `makroud:"column:id,pk"`
	Slug      string        `makroud:"column:value"`
	CurrentID sql.NullInt64 `makroud:"column:current_id,fk:town_slug"`
	CreatedAt time.Time     `makroud:"column:created_at"`
	UpdatedAt pq.NullTime   `makroud:"column:updated_at"`
}

// TableName implements Model interface.
func (Slug) TableName() string {
	return "town_slug"
}
