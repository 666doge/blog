package utils

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type RespData struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func RespOk(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, RespData{
		Code: http.StatusOK,
		Message: "success",
		Data: data,
	})
}

func RespError(c *gin.Context, code int, errMsg string ) {
	c.JSON(code, RespData{
		Code: code,
		Message: errMsg,
		Data: nil,
	})
}