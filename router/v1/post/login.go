package post

import (
	"github.com/gin-gonic/gin"
	"github.com/kjor99/golesson/dao"
	"github.com/kjor99/golesson/utils"
)

func Login(c *gin.Context) {

	var userInfo utils.UserInfo

	var res utils.Respons
	c.BindJSON(&userInfo)

	if len(userInfo.Telphone) != 11 {
		res.Code = -1
		res.Massage = "手机号码必须为11位"

		c.JSON(200, gin.H{
			"Code":    res.Code,
			"message": res.Massage,
		})
		return
	}
	if len(userInfo.Password) < 6 {
		res.Code = -1
		res.Massage = "密码为空或者少于6位数"
		c.JSON(200, gin.H{
			"Code":    res.Code,
			"message": res.Massage,
		})
		return
	}
	userInfo.Password = utils.ToMd5(userInfo.Password)
	db := dao.DB.Where("telphone=? and password=?", userInfo.Telphone, userInfo.Password).First(&userInfo)
	if db.RowsAffected == 0 {
		res.Code = -1
		res.Massage = "账号或者密码错误"
		c.JSON(200, gin.H{
			"Code":    res.Code,
			"message": res.Massage,
		})
	}
	if db.RowsAffected >= 1 {
		res.Code = 0
		res.Massage = "登录成功"
		c.JSON(200, gin.H{
			"Code":    res.Code,
			"message": res.Massage,
		})
	}

}
