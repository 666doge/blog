package db

import (
	"blog/model"
)

func InsertComment(comment *model.Comment) (commentId int64, err error) {
	sqlStr := `
		insert into comment (content, article_id, from_user, at)
		values (?,?,?,?)
	`
	result, err := DB.Exec(
		sqlStr,
		comment.Content,
		comment.ArticleId,
		comment.From,
		comment.At,
	)
	if err != nil {
		return
	}
	commentId, err = result.LastInsertId()
	return
}