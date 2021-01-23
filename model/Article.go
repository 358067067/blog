package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Ttile string `gorm:"type:varchar(100);not null" JSON:"title"`
	// Category Category `gorm:"type:" JSON:"Category"`
	Cid     int    `gorm:"type:int;not null" JSON:"cid"`
	Desc    string `gorm:"type:varchar(200)" JSON:"desc"`
	Content string `gorm:"type:longtext" JSON:"content"`
	Img     string `gorm:"type:varchar(100)" JSON:"img"`
}
