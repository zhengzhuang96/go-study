// Author: zhengzhuang
// Date: 2021-07-28 16:57:23
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-28 16:57:40
// Description: 用户模块
// FilePath: /go-study/go-web-demo/controller/user_controller.go
package controller

import (
	"fmt"
	"go-web-demo/cgo"
	"net/http"
)

type UserController struct {
}

func (uc *UserController) Router(c *cgo.RouterHandler) {
	c.POST("/userLogin", uc.userLogin)
	c.Router("/userRegister", uc.userRegister)
}

// 用户登录
func (uc *UserController) userLogin(w http.ResponseWriter, r *http.Request) {
	// if r.Method == "POST" {
	// 	fmt.Fprintf(w, "返回的数据")
	// 	fmt.Println(r.Method)
	// } else {
	// 	fmt.Fprintf(w, "请求方法错误")
	// }
	fmt.Fprintf(w, "哈哈哈")
}

// 用户注册
func (uc *UserController) userRegister(w http.ResponseWriter, r *http.Request) {
	// 1:获取用户的用户名和密码

	// 2:检测该用户是否在库中

	// 3:检测该用户密码是否正确

	// 4:登录成功返回token
}
