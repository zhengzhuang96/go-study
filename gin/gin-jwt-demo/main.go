// Author: zhengzhuang
// Date: 2021-07-27 15:25:44
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-27 15:25:52
// Description: gin-jwt-demo
// FilePath: /go-study/gin/gin-jwt-demo/main.go
package main

import (
	"gin-jwt-demo/middleware"
	"gin-jwt-demo/routers"
)

func main() {
	r := routers.SetupRouter()

	// 设置全局跨域访问
	r.Use(middleware.Cors())

	r.Run()
}
