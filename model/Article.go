package model

import (
	"blog/utils/errmsg"
	"log"

	"github.com/jinzhu/gorm"
)

//Article 文章
type Article struct {
	Category Category
	gorm.Model
	Title      string `gorm:"type:varchar(100);not null" json:"title"`
	CategoryID int    `gorm:"type:int;not null" json:"category_id"`
	Desc       string `gorm:"type:varchar(200)" json:"desc"`
	Content    string `gorm:"type:longtext" json:"content"`
	Img        string `gorm:"type:varchar(100)" json:"img"`
}

//CreateArticle 新增
func (a *Article) CreateArticle() int {
	if err := db.Create(a).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//UpdateArticle 更新
func (a *Article) UpdateArticle(id int) int {
	var maps = make(map[string]interface{})
	maps["title"] = a.Title
	maps["cid"] = a.CategoryID
	maps["desc"] = a.Desc
	maps["content"] = a.Content
	maps["img"] = a.Img
	if err := db.Model(&a).Where("id = ? ", id).Updates(maps).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//DeleteArticle 删除
func (a *Article) DeleteArticle(id int) int {
	if err := db.Where("id = ?", id).Delete(&Article{}).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetArticle 查询单个
func (a *Article) GetArticle(id int) {
	if err := db.Preload("Category").Where(" id = ?", id).Find(&a).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err)
	}
}

// GetAllArticles 分页查询
func (a *Article) GetAllArticles(pageSize int, pageNum int) ([]Article, int) {
	var as []Article
	if err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&as).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return as, errmsg.SUCCESS
}

//GetArticlesByCid 分类下所有文章
func (a *Article) GetArticlesByCid(cid int, pageSize int, pageNum int) []Article {
	var as []Article
	if err := db.Preload("Category").Where("category_id = ?", cid).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&as).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return as
}
