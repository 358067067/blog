package v1

import (
	"blog/model"
	"blog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//AddArticle 新增文章
func AddArticle(ctx *gin.Context) {
	var a model.Article
	_ = ctx.ShouldBindJSON(&a)
	if a.Title == "" || a.CategoryID == 0 || a.Content == "" {
		code = errmsg.ERROR_ARTICLE_NIL
	} else {
		code = a.CreateArticle()
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//EditArticle 修改
func EditArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var a model.Article
	_ = ctx.ShouldBindJSON(&a)
	if a.Title == "" || a.CategoryID == 0 || a.Content == "" {
		code = errmsg.ERROR_ARTICLE_NIL
	} else {
		code = a.UpdateArticle(id)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//DelArticle 删除
func DelArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var a model.Article
	code = a.DeleteArticle(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//GetAllArt 分页获取文章
func GetAllArt(ctx *gin.Context) {
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))
	var a model.Article
	as, code, total := a.GetAllArticles(pageSize, pageNum)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    as,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//GetArticle 一条记录
func GetArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var a model.Article
	a.GetArticle(id)
	code = errmsg.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    a,
		"message": errmsg.GetErrMsg(code),
	})
}

//GetArticlesByCid 按分类cid查询文章
func GetArticlesByCid(ctx *gin.Context) {
	cid, _ := strconv.Atoi(ctx.Query("cid"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))
	var a model.Article
	as, total := a.GetArticlesByCid(cid, pageSize, pageNum)
	code = errmsg.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    as,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}
