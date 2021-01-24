package model

import (
	"blog/utils/errmsg"

	"github.com/jinzhu/gorm"
)

//Category 分类
type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//CreatCategory 新增
func CreatCategory(c *Category) int {
	if err := db.Create(c).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//UpdateCategory 更新
func UpdateCategory(id int, c *Category) int {
	var maps = make(map[string]string)
	maps["name"] = c.Name
	if err := db.Model(&Category{}).Where("id = ? ", id).Updates(maps).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteCategory 删除
func DeleteCategory(id int) int {
	if err := db.Where("id = ?", id).Delete(&Category{}).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//CheckCategory 重名
func CheckCategory(name string) int {
	var c Category
	db.Select("id").Where("name = ? ", name).First(&c)
	if c.ID > 0 {
		return int(c.ID)
	}
	return errmsg.SUCCESS
}

//GetCategories 查询所有分类
func GetCategories() []Category {
	var cs []Category
	if err := db.Find(&cs).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cs
}
