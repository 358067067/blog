package model

import (
	"blog/utils/errmsg"
	"encoding/base64"
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
)

//User 结构体
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username"`
	Password string `gorm:"type:varchar(20);not null " json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

//scryptPwd 密码加密
func scryptPwd(password string) string {
	// 密码加密
	// password
	// 盐值
	salt := make([]byte, 8)
	salt = []byte{33, 22, 6, 0x58, 9, 0x6a, 5, 11}
	// 消耗cpu/memory
	// N>1
	// r*p<2^30
	// 密码长度
	const keyLen = 10
	HashPwd, err := scrypt.Key([]byte(password), salt, 1<<8, 8, 1, keyLen)
	if err != nil {
		log.Fatal(err)
	}
	rPwd := base64.StdEncoding.EncodeToString(HashPwd)
	return rPwd
}

//BeforeSave 保存前加密密码(带有事务)
func (u *User) BeforeSave(tx *gorm.DB) error {
	u.Password = scryptPwd(u.Password)
	return nil
}

//CreateUser 添加用户
func (u *User) CreateUser() int {
	// u.Password = scryptPwd(u.Password)
	if err := db.Create(u).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//CheckUser 查询用户是否存在
func (u *User) CheckUser(name string) int {
	db.Select("id").Where("username = ?", name).First(&u)
	if u.ID > 0 {
		return int(u.ID)
	}
	return errmsg.SUCCESS
}

//GetUsers 查询用户
func (u *User) GetUsers(pagaSize int, pageNum int) []User {
	var us []User
	if err := db.Limit(pagaSize).Offset((pageNum - 1) * pagaSize).Find(&us).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return us
}

//DeleteUser 删除用户
func (u *User) DeleteUser(id int) int {
	if err := db.Where("id = ? ", id).Delete(&User{}).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//UpdateUser 更新用户
func (u *User) UpdateUser(id int) int {
	var maps = make(map[string]interface{})
	maps["username"] = u.Username
	maps["role"] = u.Role
	if err := db.Model(&User{}).Where("id = ? ", id).Updates(maps).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//CheckLogin 登录验证
func (u *User) CheckLogin(username string, password string) int {
	db.Where("username = ?", username).First(&u)
	if u.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if scryptPwd(password) != u.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if u.Role != 0 {
		return errmsg.ERROR_USER_ROLE
	}
	return errmsg.SUCCESS
}
