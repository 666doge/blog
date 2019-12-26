package controller

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"net/http"
	"blog/db"
	"blog/model"
	"blog/utils"
)

func InsertComment (c *gin.Context) {
	content := c.PostForm("content")
	articleIdStr := c.PostForm("articleId")
	fromStr := c.PostForm("from")
	atStr := c.PostForm("at")

	articleId, _ := strconv.ParseInt(articleIdStr, 10, 64)
	from, _ := strconv.ParseInt(fromStr, 10, 64)
	at, _ := strconv.ParseInt(atStr, 10, 64)

	comment := &model.Comment{
		Content: content,
		ArticleId: articleId,
		From: from,
		At: at,
	}

	commentId, err := db.InsertComment(comment)
	if err != nil {
		utils.RespError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespOk(c, commentId) 
}