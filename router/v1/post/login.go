package post

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kjor99/golesson/dao"
)

func init() {
	dburl := "../golesson/conf/config.json"
	DB = dao.Conn(dburl)
}

func Login(c *gin.Context) {

	var userInfo UserInfo

	c.BindJSON(&userInfo)
	log.Default()

	if len(userInfo.Telphone) != 11 {

		c.JSON(403, "手机号码必须为11位")
		return
	}
	if len(userInfo.Password) < 6 {
		c.JSON(403, "密码为空或者少于6位数")
		return
	}
	DB.AutoMigrate(&UserInfo{})
	fmt.Print("----------")
	db := DB.Where("telphone=? and password=?", userInfo.Telphone, userInfo.Password).First(&userInfo)
	if db.RowsAffected == 0 {
		c.JSON(403, "账号或者密码错误")
	}
	if db.RowsAffected == 1 {
		c.JSON(200, "登录成功")
	}

	if len(userInfo.Username) == 0 {
		userInfo.Username = randStr(10)
		c.JSON(200, userInfo.Username)
		return
	}

}
