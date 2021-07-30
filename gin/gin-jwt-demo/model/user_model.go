// Author: zhengzhuang
// Date: 2021-07-27 16:34:08
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-27 16:34:08
// Description: In User Settings Edit
// FilePath: /go-study/gin/gin-jwt-demo/model/user.go
package model

// User 用户类
type User struct {
	Id         string `json:"userId"`
	Name       string `json:"userName"`
	Gender     string `json:"gender"`
	Phone      string `json:"userMobile"`
	Pwd        string `json:"pwd"`
	Permission string `json:"permission"`
}

// LoginReq 登录请求参数类
type LoginReq struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
