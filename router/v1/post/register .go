package post

import (
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kjor99/golesson/dao"
)

type UserInfo struct {
	Username string `gorm:"username"`
	Telphone string `gorm:"telphone"`
	Password string `gorm:"password"`
}

var DB *gorm.DB

const str = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func Register(c *gin.Context) {
	dburl := "../golesson/conf/config.json"
	DB = dao.Conn(dburl)
	var userInfo UserInfo

	c.BindJSON(&userInfo)
	defer DB.Close()

	if len(userInfo.Telphone) != 11 {

		c.JSON(403, "手机号码必须为11位")
		return
	}
	if len(userInfo.Password) < 6 {
		c.JSON(403, "密码为空或者少于6位数")
		return
	}
	if len(userInfo.Username) == 0 {
		userInfo.Username = randStr(10)
	}
	DB.AutoMigrate(&UserInfo{})
	fmt.Print("----------")
	db := DB.Where("telphone=?", userInfo.Telphone).FirstOrCreate(&userInfo)
	if db.RowsAffected == 0 {
		c.JSON(403, "此手机已注册")
	}
	if db.RowsAffected == 1 {
		c.JSON(200, "注册成功")
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
