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
func (c *Category) CreatCategory() int {
	if err := db.Create(c).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//UpdateCategory 更新
func (c *Category) UpdateCategory(id int) int {
	var maps = make(map[string]string)
	maps["name"] = c.Name
	if err := db.Model(&Category{}).Where("id = ? ", id).Updates(maps).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteCategory 删除
func (c *Category) DeleteCategory(id int) int {
	if err := db.Where("id = ?", id).Delete(&Category{}).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//CheckCategory 重名
func (c *Category) CheckCategory(name string) int {
	db.Select("id").Where("name = ? ", name).First(&c)
	if c.ID > 0 {
		return int(c.ID)
	}
	return errmsg.SUCCESS
}

//GetCategories 查询所有分类
func (c *Category) GetCategories() []Category {
	var cs []Category
	if err := db.Find(&cs).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cs
}
