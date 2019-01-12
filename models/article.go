package models

import (
	"database/sql"
	"time"

	"github.com/lib/pq"

	"github.com/EdouardParis/town/constants"
)

type Article struct {
	ID              int64          `makroud:"column:id,pk"`
	Title           string         `makroud:"column:title"`
	Subtitle        sql.NullString `makroud:"column:subtitle"`
	Body            string         `makroud:"column:body"`
	Lang            int64          `makroud:"column:lang"`
	AmountCollected int64          `makroud:"column:amount_collected"`
	AmountReceived  int64          `makroud:"column:amount_received"`
	Status          int64          `makroud:"column:status"`
	Slug            string         `makroud:"column:slug"`

	CreatedAt   time.Time   `makroud:"column:created_at"`
	UpdatedAt   pq.NullTime `makroud:"column:updated_at"`
	PublishedAt pq.NullTime `makroud:"column:published_at"`

	AddressID sql.NullInt64 `makroud:"column:address_id,fk:town_address"`
	Address   *Address      `makroud:"relation:address_id"`

	NodeID sql.NullInt64 `makroud:"column:node_id,fk:town_node"`
	Node   *Node         `makroud:"relation:node_id"`

	Reactions []Reaction
	Comments  []Comment
}

// TableName implements Model interface.
func (Article) TableName() string {
	return "town_article"
}

func (a Article) IsPublished() bool {
	return a.Status == constants.ArticleStatusPublished
}
