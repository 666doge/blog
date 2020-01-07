package model

type Comment struct {
	Id int64	`json:"id" db:"id"`
	Content string 	`json:"content" db:"content"`
	ArticleId int64 `json:"articleId" db:"article_id"` // 哪篇文章下的评论
	From int64 `json:"from" db:"from_user"` // userId, 评论由谁发出的
	At int64 `json:"at" db:"at"` // comment Id, 回复的哪一条评论
}