package controller

import (
	"strconv"
	"blog/model"
	"blog/db"
	"blog/utils"
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	summary := c.PostForm("summary")
	categoryIdStr := c.PostForm("categoryId")
	categoryId, _ := strconv.Atoi(categoryIdStr)
	username := c.PostForm("username")

	article := &model.Article{
		Title: title,
		Content: content,
		Summary: summary,
		CategoryId: categoryId,
		Username: username,
	}
	articleId, err := db.InsertArticle(article)
	if err != nil {
		utils.RespError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespOk(c, articleId)
}