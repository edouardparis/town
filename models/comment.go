package models

import "time"

type Comment struct {
	ID        int64  `makroud:"column:id,pk"`
	ArticleID int64  `makroud:"column:article_id,fk:town_article"`
	OrderID   int64  `makroud:"column:order_id,fk:town_order"`
	Body      string `makroud:"column:body"`

	CreatedAt time.Time `makroud:"column:created_at"`
}

// TableName implements Model interface.
func (Comment) TableName() string {
	return "town_article_comment"
}
