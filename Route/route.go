package Route

import (
	"dzc.com/Model"
	"dzc.com/Route/Middleware/Auth"
	"dzc.com/Route/Middleware/Exception"
	"dzc.com/Route/Middleware/Forbid"
	"dzc.com/Route/Middleware/Logger"
	"dzc.com/Route/Middleware/Operation"
	"dzc.com/Service"
	"dzc.com/Service/Mysql"
	"dzc.com/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetRoute(engine *gin.Engine) {

	// 配置中间件
	setMiddleware(engine)
	// 账号相关
	accountGroup := engine.Group("/account")
	{
		// 注册
		accountGroup.POST("/register", func(context *gin.Context) {
			var registerMod Model.RegisterBindModel
			if err := context.ShouldBind(&registerMod); err != nil {
				context.JSON(500, Utils.NewResultError500(err.Error()))
				context.Abort()
				return
			}
			if Service.ShareUserService().NameExist(registerMod.Name) {
				context.JSON(500, Utils.NewResultError500("已存在注册用户"))
				context.Abort()
				return
			}
			// &{dzc c4ca4238a0b923820dcc509a6f75849b d4b83f9ab7105c90278cfca938b77a2f 2020-05-05 02:06:27 2020-05-05 02:06:27 362138}
			context.JSON(200, Utils.NewResultSuccess200(Service.ShareUserService().Regist(registerMod.Name, registerMod.Password)))
		})

		// 登录
		accountGroup.GET("/login", func(context *gin.Context) {
			var loginMod Model.RegisterBindModel
			err := context.ShouldBindQuery(&loginMod)
			if err != nil {
				context.JSON(500, Utils.NewResultError500(err.Error()))
				context.Abort()
				return
			}

			if !(Service.ShareUserService().NameExist(loginMod.Name)) {
				context.JSON(500, Utils.NewResultError500("用户不存在"))
				context.Abort()
				return
			}

			UserModel := Service.ShareUserService().LoginUserInfo(loginMod.Name, loginMod.Password)
			if UserModel.Id == 0 {
				context.JSON(500, Utils.NewResultError500("密码错误"))
				context.Abort()
				return
			}
			context.JSON(200, Utils.NewResultSuccess200(UserModel.Token))
			//var dbs []Model.Databases
			//AppInit.GetDBCon().Raw("show databases").Scan(&dbs).Debug()
			//fmt.Print(dbs)
			//context.JSON(200, Utils.NewResultSuccess200(dbs))
		})
	}

	// 数据库连接类
	connectsGroup := engine.Group("/connects")
	{
		// 新增数据库连接信息
		connectsGroup.POST("add", func(context *gin.Context) {
			var insertMod Model.InsertConnects
			if err := context.ShouldBind(&insertMod); err != nil {
				context.JSON(500, Utils.NewResultError500(err.Error()))
				context.Abort()
				return
			}
			uuid := Utils.UUid()
			insertMod.Uuid = uuid
			//fmt.Print(insertMod)
			//context.Abort()
			//return
			_, err := Service.ShareDbsService().SetDb(uuid, insertMod)
			if err != nil {
				context.JSON(500, Utils.NewResultError500(err.Error()))
				context.Abort()
				return
			}
			userInfo, _ := Utils.GetAuthUser(context)
			insertMod.Uid = userInfo.Id
			insertRes, msg := Service.ShareConnectsService().Add(insertMod)
			if len(msg) > 0 {
				context.JSON(500, Utils.NewResultError500(msg))
				context.Abort()
				return
			}
			context.JSON(200, Utils.NewResultSuccess200(insertRes))
		})

		// 编辑 TODO
		// 删除 TODO

		// 列表
		connectsGroup.GET("/", func(context *gin.Context) {
			userInfo, _ := Utils.GetAuthUser(context)
			uid := userInfo.Id
			page := context.DefaultQuery("page", "1")
			pageNum := context.DefaultQuery("pageNum", "10")
			var lists []Model.Connects
			lists = Service.ShareConnectsService().List(uid, Utils.StringToint(page), Utils.StringToint(pageNum))
			count := Service.ShareConnectsService().Total(uid)
			context.JSON(200, Utils.NewResultSuccess200(gin.H{
				"list":  lists,
				"total": count,
			}))
		})
	}

	// 数据库连接操作类
	operationsGroup := engine.Group("/operation").Use(Operation.SetUp())
	{
		operationsGroup.GET("databases", func(context *gin.Context) {
			var selctMod Model.SelectDbs
			if err := context.ShouldBind(&selctMod); err != nil {
				context.JSON(500, Utils.NewResultError500(err.Error()))
				context.Abort()
				return
			}
			res, err := Mysql.ShareSelectService(context).ShowDatabases(selctMod)
			if err != nil {
				context.JSON(500, Utils.NewResultError500(err.Error()))
				context.Abort()
				return
			}
			context.JSON(200, Utils.NewResultSuccess200(res))
		})

		operationsGroup.GET("tablesInfo", func(context *gin.Context) {

			dbName, exist := context.GetQuery("db_name")
			if !exist {
				context.JSON(500, Utils.NewResultError500("数据库名称未指定"))
				context.Abort()
				return
			}

			res, err := Mysql.ShareSelectService(context).ShowTables(dbName)
			if err != nil {
				context.JSON(500, Utils.NewResultError500(err.Error()))
				context.Abort()
				return
			}
			context.JSON(200, Utils.NewResultSuccess200(res))
		})

		operationsGroup.GET("tableInfo", func(context *gin.Context) {

			var selectTableInfo Model.SelectTableInfo
			err := context.ShouldBindQuery(&selectTableInfo)
			if err != nil {
				context.JSON(500, Utils.NewResultError500(err.Error()))
				context.Abort()
				return
			}

			res, err := Mysql.ShareSelectService(context).ShowTableInfo(selectTableInfo.DbName, selectTableInfo.TableName)
			if err != nil {
				context.JSON(500, Utils.NewResultError500(err.Error()))
				context.Abort()
				return
			}
			context.JSON(200, Utils.NewResultSuccess200(res))
		})
	}

	// not route response
	setNoRouteResponse(engine)

	// ping response
	setTestRoute(engine)
}

func setNoRouteResponse(engine *gin.Engine) {
	engine.NoRoute(func(context *gin.Context) {
		context.JSON(404, Utils.NewResultError404("没有路由对应"))
	})
	engine.NoMethod(func(context *gin.Context) {
		context.JSON(405, Utils.NewResultError405("没有方法对应"))
	})
}

func setTestRoute(engine *gin.Engine) {
	engine.GET("/ping", func(context *gin.Context) {
		userModel, exist := Utils.GetAuthUser(context)
		if exist {
			fmt.Print(userModel)
		}
		context.JSON(200, Utils.NewResultSuccess200("pong"))
	})
}

func setMiddleware(engine *gin.Engine) {
	engine.Use(Logger.SetUp(), Forbid.SetUp(), Auth.SetUp(), Exception.SetUp())
	//engine.Use(Logger.SetUp(),Exception.SetUp())
}
