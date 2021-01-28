package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	"blog/utils/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var code int

//AddUser 添加用户
func AddUser(c *gin.Context) {
	var u model.User
	//将请求中的json绑定到u中
	_ = c.ShouldBindJSON(&u)
	msg, code := validator.Validator(u)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	if u.Username != "" && u.Password != "" {
		code = u.CheckUser(u.Username)
		if code == errmsg.SUCCESS {
			code = u.CreateUser()
		} else {
			code = errmsg.ERROR_USERNAME_USED
			c.Abort()
		}
	} else {
		code = errmsg.ERROR
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    u,
		"message": errmsg.GetErrMsg(code),
	})
}

//GetUsers 获取用户列表
func GetUsers(c *gin.Context) {
	var u model.User
	//Query获取Get请求参数
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0 && pageNum == 0 {
		pageSize = -1
	}
	us, total := u.GetUsers(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    us,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//EditUser 编辑用户
func EditUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var u model.User
	_ = c.ShouldBindJSON(&u)
	if u.Username != "" {
		if code = u.CheckUser(u.Username); code > 0 && code != id && code != errmsg.SUCCESS {
			code = errmsg.ERROR_USERNAME_USED
			// Abort 中断后续所有函数执行（包括即将到来的中间件链式调用）
			// next 执行下一个中间件的链式调用
			// 如果没有next，会先执行中间件，再执行业务方法
			c.Abort()
		} else {
			code = u.UpdateUser(id)
		}
	} else {
		code = errmsg.ERROR
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    u,
		"message": errmsg.GetErrMsg(code),
	})
}

//DelUser 删除用户
func DelUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var u model.User
	code := u.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//UserExist 删除用户
func UserExist(c *gin.Context) {
}
