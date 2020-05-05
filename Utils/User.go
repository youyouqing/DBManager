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
