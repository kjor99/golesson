package post

type UserInfo struct {
	Username string `gorm:"username"`
	Telphone string `gorm:"telphone"`
	Password string `gorm:"password"`
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
