package main
import (
	"github.com/gin-gonic/gin"
	"blog/controller"
	"blog/db"
)

func main() {
	initDb()

	startService()
}

func initDb() {
	dns := "root:xsN231564@tcp(localhost:3306)/blogger?parseTime=true&loc=Local"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
}

func startService() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/user/list", controller.GetUserList)
		v1.POST("/user", controller.AddUser)
		v1.GET("/user", controller.GetUserById)
		v1.DELETE("/user", controller.DeleteUserById)

		v1.GET("/article/list", controller.GetArticleList)
		v1.GET("/article", controller.GetArticleById)
		v1.POST("/article", controller.CreateArticle)
		v1.GET("/article/getByCateId", controller.GetArticleByCagegoryId)

		v1.GET("/category/list", controller.GetCategoryList)
	}

	router.GET("/ping", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":9527")
}