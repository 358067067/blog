package v1

import (
	"blog/middleware"
	"blog/model"
	"blog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Login 登录
func Login(ctx *gin.Context) {
	var u model.User
	var code int
	var token string
	ctx.ShouldBindJSON(&u)
	code = u.CheckLogin(u.Username, u.Password)
	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(u.Username)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
