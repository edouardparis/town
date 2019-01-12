package models

import (
	"time"
)

type ArticleReaction struct {
	ID        int64  `makroud:"column:id,pk"`
	ArticleID int64  `makroud:"column:article_id,fk:town_article"`
	OrderID   int64  `makroud:"column:order_id"`
	Emoji     string `makroud:"column:emoji"`

	CreatedAt time.Time `makroud:"column:created_at"`
}

// TableName implements Model interface.
func (ArticleReaction) TableName() string {
	return "town_article_reaction"
}
