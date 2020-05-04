package Route

import (
	"dzc.com/Route/Middleware/Logger"
	"dzc.com/Utils"
	"github.com/gin-gonic/gin"
)

func SetRoute(engine *gin.Engine)  {

	// 配置中间件
	setMiddleware(engine)

	//
	accountGroup := engine.Group("/account")
	{
		// 账号相关接口
		accountGroup.GET("/users", func(context *gin.Context) {
			context.JSON(200,Utils.NewResultSuccess200("测试数据"))
		})
	}
	//EntriGroup := gee.Group("/task")
	//{
	//	// 企业相关接口
	//	taskApiGroup.POST("/add", HandleTaskAdd)
	//	taskApiGroup.POST("/addAndStart", HandleTaskAddAndStart)
	//	taskApiGroup.GET("/start/:id", HandleTaskStart)
	//	taskApiGroup.GET("/stop/:id", HandleTaskStop)
	//}
}

func setMiddleware(engine *gin.Engine)  {
	engine.Use(Logger.SetLogger())
}
