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

func GetArticleList(c *gin.Context) {
	list, err := db.GetArticleList()
	if err != nil {
		utils.RespError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespOk(c, list)
}


func GetArticleById(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	article, err := db.GetArticleById(id)
	if err != nil {
		utils.RespError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespOk(c, article)
}

func GetArticleByCagegoryId(c *gin.Context) {
	cateIdStr := c.Query("cateId")
	id, _ := strconv.ParseInt(cateIdStr, 10, 64)
	list, err := db.GetArticleByCagegoryId(id)
	if err != nil {
		utils.RespError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespOk(c, list)
}