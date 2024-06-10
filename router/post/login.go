package post

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `gorm:"username"`
	Telphone string `gorm:"telphone"`
	Password string `gorm:"password"`
}

func Login(c *gin.Context) {

	var userInfo UserInfo

	c.BindJSON(&userInfo)
	fmt.Printf("userInfo: %v\n", userInfo)

	fmt.Printf("userInfo.password: %v\n", userInfo.Password)
	if len(userInfo.Password) < 6 {
		c.JSON(200, "密码为空或者少于6位数")
	}
	if userInfo.Username == "" {
		c.String(200, "用户名为空", userInfo.Username)
	}

}
