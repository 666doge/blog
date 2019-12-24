package controller
import (
	"blog/db"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCategoryList(c *gin.Context) {
	list, err := db.GetCategoryList()
	if err != nil {
		utils.RespError(c, http.StatusInternalServerError, err.Error())
	}
	utils.RespOk(c, list)
}
