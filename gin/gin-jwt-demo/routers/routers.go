// Author: zhengzhuang
// Date: 2021-07-27 15:27:56
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-27 15:27:56
// Description: In User Settings Edit
// FilePath: /go-study/gin/gin-jwt-demo/routers/routers.go
package routers

import (
	"gin-jwt-demo/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	app := gin.Default()

	// app.Use(middleware.JWTAuth())

	new(controller.UserController).Router(app)

	return app
}
