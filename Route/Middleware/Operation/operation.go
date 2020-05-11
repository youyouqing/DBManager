package Operation

import (
	"dzc.com/Service"
	"dzc.com/Utils"
	"github.com/gin-gonic/gin"
)

func SetUp() gin.HandlerFunc {

	return func(context *gin.Context) {
		userInfo, _ := Utils.GetAuthUser(context)
		uuid, _ := context.GetQuery("uuid")
		if len(uuid) > 0 {
			mod := Service.ShareConnectsService().Detial(uuid)
			if mod.Id > 0 && mod.Uid == userInfo.Id {
				context.Set("connInfo", mod)
			} else {
				context.JSON(403, Utils.NewResultError403("数据库连接信息和用户不对应"))
				context.Abort()
				return
			}
		}
		context.Next()
	}
}
