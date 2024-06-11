package post

import (
	"time"
)

type UserInfo struct {
	Id        uint64    `gorm:"id;primaryKey"`
	Username  string    `gorm:"username"`
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

type Respons struct {
	code    int    `gorm:"code"`
	massage string `gorm:"massage"`
}
