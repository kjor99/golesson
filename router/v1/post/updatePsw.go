package post

import (
	"github.com/gin-gonic/gin"
	"github.com/kjor99/golesson/dao"
	"github.com/kjor99/golesson/utils"
)

func UpdatePsw(c *gin.Context) {

	DB = dao.Conn()
	var updateUserInfo UpdateUserInfo
	var res Respons
	c.BindJSON(&updateUserInfo)
	defer DB.Close()

	if len(updateUserInfo.Telphone) != 11 {
		res.code = -1
		res.massage = "手机号码必须为11位"
		c.JSON(200, gin.H{
			"code":    res.code,
			"message": res.massage,
		})

		return
	}
	if len(updateUserInfo.Password) < 6 {
		res.code = -1
		res.massage = "密码为空或者少于6位数"
		c.JSON(200, gin.H{
			"code":    res.code,
			"message": res.massage,
		})
		return
	}
	updateUserInfo.NewPassword = utils.ToMd5(updateUserInfo.NewPassword)
	updateUserInfo.Password = utils.ToMd5(updateUserInfo.Password)

	db := DB.Where("telphone=? and password=?", updateUserInfo.Telphone, updateUserInfo.Password).First(&UserInfo{})
	if db.RowsAffected == 0 {
		res.code = -1
		res.massage = "账号或者密码错误"
		c.JSON(200, gin.H{
			"code":    res.code,
			"message": res.massage,
		})
	}
	//当密码与账号匹配时才能修改密码
	if db.RowsAffected == 1 {
		db = DB.Model(&UserInfo{}).Updates(UserInfo{Password: updateUserInfo.NewPassword})
		if db.RowsAffected == 1 {
			res.code = 0
			res.massage = "密码修改成功"
			c.JSON(200, gin.H{
				"code":    res.code,
				"message": res.massage,
			})

		} else {

			res.code = -1
			res.massage = "密码修改失败"
			c.JSON(200, gin.H{
				"code":    res.code,
				"message": res.massage,
			})
		}

	}

}
