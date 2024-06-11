package post

import (
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/kjor99/golesson/dao"
	"github.com/kjor99/golesson/utils"
)

const str = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func Register(c *gin.Context) {

	dao.DB = dao.Conn()
	var userInfo UserInfo
	var res Respons

	c.BindJSON(&userInfo)
	defer dao.DB.Close()

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
	if len(userInfo.Username) == 0 {
		userInfo.Username = randStr(10)
	}

	dao.DB.AutoMigrate(&userInfo)
	//手机号码加密
	userInfo.Password = utils.ToMd5(userInfo.Password)

	db := dao.DB.Where("telphone=?", userInfo.Telphone).FirstOrCreate(&userInfo)
	if db.RowsAffected == 0 {
		res.code = -1
		res.massage = "该手机已注册"

		c.JSON(200, gin.H{
			"code":    res.code,
			"message": res.massage,
		})
	}
	if db.RowsAffected == 1 {
		res.code = 0
		res.massage = "注册成功"
		c.JSON(200, gin.H{
			"code":     res.code,
			"masssage": res.massage,
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
