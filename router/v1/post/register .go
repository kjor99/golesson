package post

import (
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/kjor99/golesson/dao"
	"github.com/kjor99/golesson/utils"
)

const str = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func Register(c *gin.Context) {

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
	if len(userInfo.Username) == 0 {
		userInfo.Username = randStr(10)
	}

	dao.DB.AutoMigrate(&userInfo)
	//手机号码加密
	userInfo.Password = utils.ToMd5(userInfo.Password)

	db := dao.DB.Where("telphone=?", userInfo.Telphone).FirstOrCreate(&userInfo)
	if db.RowsAffected == 0 {
		res.Code = -1
		res.Massage = "该手机已注册"

		c.JSON(200, gin.H{
			"Code":    res.Code,
			"message": res.Massage,
		})
	}
	if db.RowsAffected == 1 {
		res.Code = 0
		res.Massage = "注册成功"
		c.JSON(200, gin.H{
			"Code":     res.Code,
			"masssage": res.Massage,
		})
	}

}

func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = str[rand.Int63()%int64(len(str))]
	}
	return string(b)

}
