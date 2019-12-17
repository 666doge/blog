package db

import (
	// "fmt"
	// "time"
	"blog/model"
)

func InsertArticle(article *model.Article) (articleId int64, err error) {
	sqlStr := `insert into
				article (title, content, summary, category_id, username, view_count, comment_count, status)
				values(?, ?, ?, ?, ?, ?, ?, ?)
			`
	result, err := DB.Exec(
		sqlStr,
		article.Title,
		article.Content,
		article.Summary,
		article.CategoryId,
		article.Username,
		article.ViewCount,
		article.CommentCount,
		article.Status,
	)
	if err != nil {
		return
	}
	
	articleId, err = result.LastInsertId()
	return
}