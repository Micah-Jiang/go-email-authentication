package model

import (
	"fmt"
	"go-email-authentication/pkg/config"
	"gorm.io/gorm"
	"reflect"
	"time"
)

type SysUser struct {
	id          int       `gorm:"primaryKey"`
	Username    string    `gorm:"index" json:"username" binding:"required"`
	Password    string    `json:"password" binding:"required"`
	Email       string    `gorm:"index" json:"email" binding:"required"`
	CreatedTime time.Time `gorm:"autoCreateTime"`
}

var db *gorm.DB

func init() {
	d := config.GetDB()
	db = d
}

func (u SysUser) TableName() string {
	return "sys_user"
}

func (u SysUser) IsEmpty() bool {
	return reflect.DeepEqual(u, SysUser{})
}

func InsertUser(user SysUser) {
	db.Create(user)
}

func QueryUserCountByEmail(email string) int64 {
	var count int64
	db.Where("email = ?", email).Count(&count)
	fmt.Printf("email register count: %d\n", count)
	return count
}

func QueryUserCountByUserName(username string) int64 {
	var count int64
	db.Where("username = ?", username).Count(&count)
	return count
}

func QueryUserByUserNameOrEmail(username, email string) SysUser {
	user := SysUser{}
	db.Where("username = ? or email = ?", username, email).Find(&user)
	return user
}
