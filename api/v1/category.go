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
		if code = c.CheckCategory(c.Name); code == errmsg.SUCCESS {
			code = c.CreatCategory()
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

//EditCategory 修改分类
func EditCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var c model.Category
	_ = ctx.ShouldBindJSON(&c)
	if c.Name != "" {
		if code = c.CheckCategory(c.Name); code > 0 && code != id && code != errmsg.SUCCESS {
			code = errmsg.ERROR_CNAME_USED
			ctx.Abort()
		} else {
			code = c.UpdateCategory(id)
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
	var c model.Category
	var cs []model.Category
	cs = c.GetCategories()
	code = errmsg.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    cs,
		"message": errmsg.GetErrMsg(code),
	})
}

//DelCategory 删除分类
func DelCategory(ctx *gin.Context) {
	var c model.Category
	id, _ := strconv.Atoi(ctx.Param("id"))
	c.DeleteCategory(id)
	code = errmsg.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
