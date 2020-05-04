package AppInit

import (
	"dzc.com/Utils"
	"github.com/gin-gonic/gin"
)

type ginEngine struct {
	*gin.Engine
	ServerPort string
}

func GetGin(isProduct bool) *ginEngine {
	return &ginEngine{engine,Utils.ShareConfigInstance(isProduct).GetConfigFromKey("server_port")}
}

func (this *ginEngine)Start()  {

	GetDB()

	this.Run(this.ServerPort)
}

var engine *gin.Engine

func init() {
	engine = gin.Default()
}

