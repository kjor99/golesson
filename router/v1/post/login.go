package post

import (
	"github.com/gin-gonic/gin"
	"github.com/kjor99/golesson/dao"
)

func Login(c *gin.Context) {

	DB = dao.Conn()

	var userInfo UserInfo

	var res Respons
	c.BindJSON(&userInfo)
	defer DB.Close()

	if len(userInfo.Telphone) != 11 {
		res.code = -1
		res.massage = "手机号码必须为11位"

		c.JSON(200, gin.H{
			"code":    res.code,
			"message": res.massage,
		})
		return
	}
	if len(userInfo.Password) < 6 {
		res.code = -1
		res.massage = "密码为空或者少于6位数"
		c.JSON(200, gin.H{
			"code":    res.code,
			"message": res.massage,
		})
		return
	}
	DB.AutoMigrate(&UserInfo{})
	db := DB.Where("telphone=? and password=?", userInfo.Telphone, userInfo.Password).First(&userInfo)
	if db.RowsAffected == 0 {
		res.code = -1
		res.massage = "账号或者密码错误"
		c.JSON(200, gin.H{
			"code":    res.code,
			"message": res.massage,
		})
	}
	if db.RowsAffected == 1 {
		res.code = 0
		res.massage = "登录成功"
		c.JSON(200, gin.H{
			"code":    res.code,
			"message": res.massage,
		})
	}

}
