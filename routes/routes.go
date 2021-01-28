package routes

import (
	v1 "blog/api/v1"
	"blog/middleware"
	"blog/utils"

	"github.com/gin-gonic/gin"
)

//InitRouter 初始化路由
func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	//引入中间件
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtTokenMid())
	{
		auth.GET("user/list", v1.GetUsers)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DelUser)

		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DelCategory)

		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DelArticle)
		//上传
		auth.POST("upload", v1.Upload)
	}

	r1 := r.Group("api/v1/")
	{
		r1.POST("user/add", v1.AddUser)
		r1.GET("category/list", v1.GetCategories)
		r1.GET("article/list", v1.GetAllArt)
		r1.GET("article/single/:id", v1.GetArticle)
		r1.GET("article/c/", v1.GetArticlesByCid)
		r1.POST("login", v1.Login)
	}
	r.Run(utils.HttpPort)
}
