package Forbid

import (
	"dzc.com/Utils"
	"github.com/gin-gonic/gin"
)

func SetUp() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Request.RequestURI == "/favicon.ico" {
			context.JSON(403, Utils.NewResultError403("谷歌浏览器图标请求拦截"))
			context.Abort()
			return
		}
		context.Next()
	}
}
