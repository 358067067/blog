package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " JSON:"username"`
	Password string `gorm:"type:varchar(20);not null " JSON:"password"`
	Role     int    `gorm:"type:int" JSON:"role"`
}
