package model

import (
	"time"
)

type Article struct {
	Id int64 `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	Content string `db:"content" json:"content"`
	Summary string `db:"summary" json:"summary"`
	ViewCount int `db:"view_count" json:"viewCount"`
	CreateTime time.Time `db:"create_time" json:"createTime"`
	UpdateTime time.Time `db:"update_time" json:"updateTime"`
	CommentCount int `db:"comment_count" json:"commentCount"`
	Username string `db:"username" json:"username"`
	CategoryId int `db:"category_id" json:"categoryId"`
	Status int `db:"status" json:"status"`
}