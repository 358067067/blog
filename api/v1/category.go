package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//AddCategory 新增分类
func AddCategory(ctx *gin.Context) {
	var c model.Category
	_ = ctx.ShouldBindJSON(&c)
	if c.Name != "" {
		if code = model.CheckCategory(c.Name); code == errmsg.SUCCESS {
			code = model.CreatCategory(&c)
		} else {
			code = errmsg.ERROR_CNAME_USED
			ctx.Abort()
		}
	} else {
		code = errmsg.ERROR
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    c,
		"message": errmsg.GetErrMsg(code),
	})
}

//UpdCategory 修改分类
func UpdCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var c model.Category
	_ = ctx.ShouldBindJSON(&c)
	if c.Name != "" {
		if code = model.CheckCategory(c.Name); code > 0 && code != id && code != errmsg.SUCCESS {
			code = model.UpdateCategory(id, &c)
		} else {
			code = errmsg.ERROR_CNAME_USED
		}
	} else {
		code = errmsg.ERROR
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    c,
		"message": errmsg.GetErrMsg(code),
	})
}

//GetCategories 获取所有分类
func GetCategories(ctx *gin.Context) {
	var cs []model.Category
	cs = model.GetCategories()
	code = errmsg.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    cs,
		"message": errmsg.GetErrMsg(code),
	})
}

//DelCategory 删除分类
func DelCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	model.DeleteCategory(id)
	code = errmsg.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
