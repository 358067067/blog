package routes

import (
	v1 "blog/api/v1"
	"blog/utils"

	"github.com/gin-gonic/gin"
)

//InitRouter 初始化路由
func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	r1 := r.Group("api/v1/user")
	{
		// User模块路由接口
		r1.POST("/add", v1.AddUser)
		r1.GET("/list", v1.GetUsers)
		r1.PUT("/:id", v1.EditUser)
		r1.DELETE("/:id", v1.DelUser)
	}
	r2 := r.Group("api/v1/category")
	{
		r2.POST("/add", v1.AddCategory)
		r2.GET("/list", v1.GetCategories)
		r2.PUT("/:id", v1.EditCategory)
		r2.DELETE("/:id", v1.DelCategory)
	}
	r3 := r.Group("api/v1/article")
	{
		r3.POST("/add", v1.AddArticle)
		r3.GET("/list", v1.GetAllArt)
		r3.GET("/single/:id", v1.GetArticle)
		r3.GET("/c/", v1.GetArticlesByCid)
		r3.PUT("/:id", v1.EditArticle)
		r3.DELETE("/:id", v1.DelArticle)
	}
	r.Run(utils.HttpPort)
}
