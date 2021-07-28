// Author: zhengzhuang
// Date: 2021-07-27 15:32:03
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-27 15:32:03
// Description: In User Settings Edit
// FilePath: /go-study/gin/gin-jwt-demo/controller/user_controller.go
package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) Router(engine *gin.Engine) {
	engine.GET("/", uc.getUserInfo)
	engine.POST("/userLogin", uc.userLogin)
}

func (uc *UserController) getUserInfo(c *gin.Context) {
	println("获取用户信息")
	c.JSON(200, map[string]interface{}{
		"name": "123",
	})
}

// Login 登录
func (uc *UserController) userLogin(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	println(n)
	// // 1.声明接收的变量
	// var json model.LoginReq
	// // 将request的body中的数据，自动按照json格式解析到结构体
	// if err := c.ShouldBind(&json); err != nil {
	// 	// 返回错误信息
	// 	// gin.H封装了生成json数据的工具
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// println(json.Mobile)
	// // 判断用户名密码是否正确
	// if json.Mobile != "18954299351" || json.Password != "123123" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"status": "200"})
	// println("用户登录")
	// var loginReq model.LoginReq
	// fmt.Println(c.Request.Body)
	// c.BindJSON(&loginReq)
	// fmt.Println(loginReq.Mobile)
	// var loginReq model.LoginReq
	// c.BindJSON(&loginReq)
	// fmt.Println(loginReq)
	// if c.BindJSON(&loginReq) != nil {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status": -1,
	// 		"msg":    "json解析失败",
	// 	})
	// 	return
	// }
	// c.JSON(200, map[string]interface{}{
	// 	"user": "123123",
	// })
	// var loginReq model.LoginReq
	// if c.BindJSON(&loginReq) == nil {
	// 	println(loginReq.Mobile)
	// 	// isPass, user, err := model.LoginCheck(loginReq)
	// 	// if isPass {
	// 	// 	generateToken(c, user)
	// 	// } else {
	// 	// 	c.JSON(http.StatusOK, gin.H{
	// 	// 		"status": -1,
	// 	// 		"msg":    "验证失败," + err.Error(),
	// 	// 	})
	// 	// }
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status": -1,
	// 		"msg":    "json 解析失败",
	// 	})
	// }
}
