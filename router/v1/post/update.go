package post

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kjor99/golesson/dao"
)

type UpdateUserInfo struct {
	UserInfo
	NewPassword string `gorm:"NewPassword"`
}

func UpdateInfo(c *gin.Context) {

	DB = dao.Conn()
	var userInfo UserInfo
	var updateUserInfo UpdateUserInfo

	c.BindJSON(&updateUserInfo)
	defer DB.Close()

	if len(updateUserInfo.Telphone) != 11 {

		c.JSON(403, "手机号码必须为11位")
		return
	}
	if len(updateUserInfo.Password) < 6 {
		c.JSON(403, "密码为空或者少于6位数")
		return
	}
	fmt.Print("----------")
	db := DB.Where("telphone=? and password=?", updateUserInfo.Telphone, updateUserInfo.Password).First(&userInfo)
	if db.RowsAffected == 0 {
		c.JSON(403, "账号或者密码错误")
	}
	if db.RowsAffected == 1 {

		db = DB.Model(&userInfo).Update("password", updateUserInfo.NewPassword)
		if db.RowsAffected == 1 {
			c.JSON(200, "密码修改成功")

		} else {
			c.JSON(403, "密码修改失败")
		}

	}

}
