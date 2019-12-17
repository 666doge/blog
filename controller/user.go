package controller

import (
	"strconv"
	"blog/model"
	"blog/db"
	"blog/utils"
	"net/http"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {

	name := c.PostForm("name")
	ageStr := c.PostForm("age")
	age, _ := strconv.ParseInt(ageStr, 10, 64)
	sexStr := c.PostForm("sex")
	sex, _ := strconv.Atoi(sexStr)

	user := &model.User{
		Name: name,
		Age: int(age),
		Sex: model.Sex(sex),
	}
	userId, err := db.InsertUser(user)
	if err != nil {
		utils.RespError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespOk(c, userId)
}

func GetUserById(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	user, err := db.GetUserById(id)
	if err != nil {
		utils.RespError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespOk(c, user)
}

func GetUserList(c *gin.Context) {
	userList, err := db.GetUserList()
	if err != nil {
		utils.RespError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespOk(c, userList)
}

func DeleteUserById(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	err := db.DeleteUserById(id)
	if err != nil {
		utils.RespError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespOk(c, nil)
}
