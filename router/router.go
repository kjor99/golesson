package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kjor99/golesson/readConf"
	"github.com/kjor99/golesson/router/v1/post"
)

func StartGin() {
	GinConfUrl := "../golesson/conf/config.json"
	r := gin.Default()
	v1 := r.Group("/v1")

	{
		v1.POST("/login", post.Login)

	}
	v2 := r.Group("/v2")
	{
		v2.GET("/login")
	}
	port := readConf.GetGinConf(GinConfUrl)
	r.Run(":" + port)

}
