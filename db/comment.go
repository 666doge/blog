package db

import (
	"blog/model"
)

func InsertComment(comment *model.Comment) (commentId int64, err error) {
	tx, err := DB.Beginx()
	if err != nil {
		return
	}
	defer func () {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	// 1.插入评论
	sqlStr := `
		insert into comment (content, article_id, from_user, at)
		values (?,?,?,?)
	`
	result, err := tx.Exec(
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

	// 2.更新文章的 comment_count字段
	sqlStr = `update article 
				set comment_count = comment_count + 1
				where id = ?
			`
	_, err = tx.Exec(sqlStr, comment.ArticleId)
	if err != nil {
		return
	}

	err = tx.Commit();

	return
}

func GetCommentList(articleId int64) (list []*model.Comment, err error) {
	sqlStr := `
		select 
			id, content, from_user, at, article_id
		from comment where article_id = ?
	`
	err = DB.Select(&list, sqlStr, articleId)
	return
}