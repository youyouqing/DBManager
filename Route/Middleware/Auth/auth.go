package Auth

import (
	"dzc.com/Service"
	"dzc.com/Utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func SetUp() gin.HandlerFunc {

	return func(context *gin.Context) {
		var token string
		if strings.Index(context.Request.RequestURI, "/account") != 0 {
			if context.Request.Header["token"] != nil {
				token = context.Request.Header.Get("token")
			}
			if len(context.Query("token")) > 0 {
				token = context.Query("token")
			}
			postToken, _ := context.GetPostForm("token")
			if len(postToken) > 0 {
				token = postToken
			}
			if len(token) == 0 {
				context.JSON(403, Utils.NewResultError403("token不能为空"))
				context.Abort()
				return
			}

			userModel, finded := Service.ShareUserService().UserInfoFromToken(token)
			if !finded {
				context.JSON(403, Utils.NewResultError403("token失效，请重新登录"))
				context.Abort()
				return
			}
			context.Set("user", userModel)
		}

		context.Next()
	}
}
