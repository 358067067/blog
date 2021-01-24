package model

import (
	"blog/utils"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

var (
	db  *gorm.DB
	err error
)

//InitDb 数据库连接
func InitDb() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPwd,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
	}
	// 禁用表明复数形式，如结构体'User'，表明是'user'，不禁用是'users'
	db.SingularTable(true)
	// 以下结构体会自动创建表
	db.AutoMigrate(&User{}, &Category{}, &Article{})
	// 空闲连接数
	db.DB().SetMaxIdleConns(10)
	// 最大连接数
	db.DB().SetMaxOpenConns(100)
	// 连接可复用时间
	db.DB().SetConnMaxLifetime(10 * time.Second)
	// db.Close()
}
