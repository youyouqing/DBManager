package main

import (
	"dzc.com/AppInit"
	"dzc.com/Route"
)

func main() {

	gee := AppInit.GetGin()

	// 设置路由 和中间件
	Route.SetRoute(gee.Engine)

	gee.Start()
}
