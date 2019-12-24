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

func GetArticleList() (articleList []*model.Article, err error) {
	sqlStr := `select
					id, title, content, summary, category_id, view_count, username, comment_count, create_time, update_time
				from article`
	err = DB.Select(&articleList, sqlStr) 
	return
}

func GetArticleById(id int64) (article *model.Article, err error) {
	article = &model.Article{}
	sqlStr := `select
				id, title, content, summary, category_id, view_count, username, comment_count, create_time, update_time
			from article
			where
				id = ?
			`
	err = DB.Get(article, sqlStr, id)
	return
}

func GetRelativeArticleList(id int64) (list []*model.Article, err error) {
	var cateId int64
	sqlStr := `select category_id from article where id = ?`
	err = DB.Get(&cateId, sqlStr, id)
	if err != nil {
		return nil, err
	}
	
	sqlStr = `select
				id, title, content, summary, category_id, view_count, username, comment_count, create_time, update_time
			from article
			where
			category_id = ? and id != ?
			`
	err = DB.Select(&list, sqlStr, cateId, id)
	return
}

func GetNextArticle(id int64) (article *model.Article, err error) {
	article = &model.Article{}
	sqlStr := `select * from article 
			where id > ?
			order by id asc
			limit 1
		`
	err = DB.Get(article, sqlStr, id)
	return
}

func GetPreArticle(id int64) (article *model.Article, err error) {
	article = &model.Article{}
	sqlStr := `select * from article 
			where id < ?
			order by id desc
			limit 1
		`
	err = DB.Get(article, sqlStr, id)
	return
}