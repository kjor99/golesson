package utils

import (
	"time"
)

type UserInfo struct {
	Id        uint64    `gorm:"id;primaryKey"`
	Username  string    `gorm:"username;size:20"`
	Telphone  string    `gorm:"telphone;primaryKey"`
	Password  string    `gorm:"password"`
	CreatedAt time.Time `gorm:"column:createtime"`
	UpdatedAt time.Time `gorm:"column:updatetime"`
}

// 修改密码的入参比注册多一个新密码参数
type UpdateUserInfo struct {
	UserInfo
	NewPassword string `gorm:"NewPassword"`
}

// 返回信息
type Respons struct {
	Code    int    `gorm:"code"`
	Massage string `gorm:"massage"`
}
