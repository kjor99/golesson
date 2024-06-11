package post

import (
	"github.com/gin-gonic/gin"
	"github.com/kjor99/golesson/dao"
	"github.com/kjor99/golesson/utils"
)

func UpdatePsw(c *gin.Context) {

	var updateUserInfo utils.UpdateUserInfo
	var res utils.Respons
	c.BindJSON(&updateUserInfo)

	if len(updateUserInfo.Telphone) != 11 {
		res.Code = -1
		res.Massage = "手机号码必须为11位"
		c.JSON(200, gin.H{
			"Code":    res.Code,
			"message": res.Massage,
		})

		return
	}
	if len(updateUserInfo.Password) < 6 {
		res.Code = -1
		res.Massage = "密码为空或者少于6位数"
		c.JSON(200, gin.H{
			"Code":    res.Code,
			"message": res.Massage,
		})
		return
	}
	updateUserInfo.NewPassword = utils.ToMd5(updateUserInfo.NewPassword)
	updateUserInfo.Password = utils.ToMd5(updateUserInfo.Password)

	db := dao.DB.Where("telphone=? and password=?", updateUserInfo.Telphone, updateUserInfo.Password).First(&utils.UserInfo{})
	if db.RowsAffected == 0 {
		res.Code = -1
		res.Massage = "账号或者密码错误"
		c.JSON(200, gin.H{
			"Code":    res.Code,
			"message": res.Massage,
		})
	}
	//当密码与账号匹配时才能修改密码
	if db.RowsAffected == 1 {
		db = dao.DB.Model(&utils.UserInfo{}).Updates(utils.UserInfo{Password: updateUserInfo.NewPassword})
		if db.RowsAffected == 1 {
			res.Code = 0
			res.Massage = "密码修改成功"
			c.JSON(200, gin.H{
				"Code":    res.Code,
				"message": res.Massage,
			})

		} else {

			res.Code = -1
			res.Massage = "密码修改失败"
			c.JSON(200, gin.H{
				"Code":    res.Code,
				"message": res.Massage,
			})
		}

	}

}
