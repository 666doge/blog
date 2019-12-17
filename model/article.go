package model

import (
	"time"
)

type Article struct {
	Title string `db:"title"`
	Content string `db:"content"`
	Summary string `db:"summary"`
	ViewCount int `db:"view_count"`
	CreateTime time.Time `db:"create_time"`
	CommentCount int `db:"comment_count"`
	Username string `db:"username"`
	CategoryId int `db:"category_id"`
	Status int `db:"status"`
}