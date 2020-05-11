package Exception

import (
	"dzc.com/Utils"
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"strings"
)

func SetUp() gin.HandlerFunc {

	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "<br>"
				}
				//fmt.Print(DebugStack)
				context.JSON(500, Utils.NewResultError500(err.(error).Error()))
				context.Abort()
				return
			}
		}()
		context.Next()
	}
}
