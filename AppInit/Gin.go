package AppInit

import (
	"dzc.com/Config"
	"github.com/gin-gonic/gin"
)

type ginEngine struct {
	*gin.Engine
	ServerPort string
}

func GetGin() *ginEngine {
	return &ginEngine{engine, Config.SERVER_PORT}
}

func (this *ginEngine) Start() {

	//GetDB()

	this.Run(this.ServerPort)
}

var engine *gin.Engine

func init() {
	engine = gin.Default()
}
