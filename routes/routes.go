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
	rv1 := r.Group("api/v1")
	{
		// User模块路由接口
		rv1.POST("user/add", v1.AddUser)
		rv1.GET("users", v1.GetUsers)
		rv1.PUT("user/:id", v1.EditUser)
		rv1.DELETE("user/:id", v1.DelUser)
		rv1.POST("category/add", v1.AddCategory)
		rv1.GET("categories", v1.GetCategories)
		rv1.PUT("category/:id", v1.EditUser)
		rv1.DELETE("category/:id", v1.DelCategory)
	}
	r.Run(utils.HttpPort)
}
