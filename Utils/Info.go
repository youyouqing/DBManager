package Utils

import (
	"dzc.com/Model"
	"github.com/gin-gonic/gin"
)

func GetAuthUser(ctx *gin.Context) (Model.Users, bool) {
	var userModel Model.Users
	userM, _ := ctx.Get("user")
	if userModel, ok := userM.(Model.Users); ok {
		return userModel, true
	}
	return userModel, false
}

func GetConnInfo(ctx *gin.Context) (Model.Connects, bool) {
	var conModel Model.Connects
	conMod, _ := ctx.Get("connInfo")
	if conMod, ok := conMod.(Model.Connects); ok {
		return conMod, true
	}
	return conModel, false
}
