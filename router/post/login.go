package post

import (
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `gorm:"username"`
	Telphone string `gorm:"telphone"`
	Password string `gorm:"password"`
}

const str = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func Login(c *gin.Context) {

	var userInfo UserInfo

	c.BindJSON(&userInfo)

	if len(userInfo.Password) < 6 {
		c.JSON(200, "密码为空或者少于6位数")
	}
	if len(userInfo.Username) == 0 {
		userInfo.Username = randStr(10)
		c.JSON(200, userInfo.Username)
	}

}

func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = str[rand.Int63()%int64(len(str))]
	}
	fmt.Printf("b: %v\n", b)
	return string(b)

}
