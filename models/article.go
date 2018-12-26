package models

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Article struct {
	ID              int64          `makroud:"column:id,pk"`
	Title           string         `makroud:"column:title"`
	Subtitle        sql.NullString `makroud:"column:subtitle"`
	BodyMD          string         `makroud:"column:body_md"`
	BodyHTML        string         `makroud:"column:body_html"`
	Lang            int64          `makroud:"column:lang"`
	AmountCollected int64          `makroud:"column:amount_collected"`
	AmountReceived  int64          `makroud:"column:amount_received"`
	Status          int64          `makroud:"column:status"`

	CreatedAt   time.Time   `makroud:"column:created_at"`
	UpdatedAt   pq.NullTime `makroud:"column:updated_at"`
	PublishedAt pq.NullTime `makroud:"column:published_at"`

	AddressID sql.NullInt64 `makroud:"column:address_id,fk:town_address"`
	Address   *Address      `makroud:"relation:address_id"`

	NodeID sql.NullInt64 `makroud:"column:node_id,fk:town_node"`
	Node   *Node         `makroud:"relation:node_id"`

	SlugID int64 `makroud:"column:slug_id,fk:town_slug"`
	Slug   *Slug `makroud:"relation:slug_id"`
}

// TableName implements Model interface.
func (Article) TableName() string {
	return "town_article"
}
