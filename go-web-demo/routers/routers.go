// Author: zhengzhuang
// Date: 2021-07-28 15:52:53
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-28 15:53:05
// Description: In User Settings Edit
// FilePath: /go-study/go-web-demo/routers/router.go
package routers

import (
	"go-web-demo/cgo"
	"go-web-demo/controller"
)

// 注册路由
func SetupRouter(handler *cgo.RouterHandler) {
	new(controller.TextController).Router(handler)
	new(controller.UserController).Router(handler)
}
