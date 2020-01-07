package controller

import (
	"fmt"
	"strconv"
	"net/http"
	"blog/db"
	"blog/model"
	"blog/utils"

	"github.com/gin-gonic/gin"
)

func InsertComment (c *gin.Context) {
	content := c.PostForm("content")
	articleIdStr := c.PostForm("articleId")
	fromStr := c.PostForm("from")
	atStr := c.PostForm("at")

	articleId, _ := strconv.ParseInt(articleIdStr, 10, 64)
	from, _ := strconv.ParseInt(fromStr, 10, 64)
	at, err := strconv.ParseInt(atStr, 10, 64)

	// 1 检查articleId是否存在
	exists, err := db.IfArticleExists(articleId)
	if exists == false {
		utils.RespError(
			c,
			http.StatusInternalServerError, 
			fmt.Sprintf("article id: %d not exists", articleId),
		)
		return;
	}
	
	if err != nil {
		utils.RespError(
			c,
			http.StatusInternalServerError, 
			fmt.Sprintf("query article id: %d failed: %s", articleId, err.Error()),
		)
		return
	}

	// 2插入评论
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

func GetCommentList(c *gin.Context) {
	articleIdStr := c.Query("articleId")
	articleId, _ := strconv.ParseInt(articleIdStr, 10, 64)
	list, err := db.GetCommentList(articleId)
	if err != nil {
		utils.RespError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespOk(c, list)
}