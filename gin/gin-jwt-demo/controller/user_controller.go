// Author: zhengzhuang
// Date: 2021-07-27 15:32:03
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-27 15:32:03
// Description: In User Settings Edit
// FilePath: /go-study/gin/gin-jwt-demo/controller/user_controller.go
package controller

import (
	"fmt"
	"gin-jwt-demo/middleware"
	"gin-jwt-demo/model"
	"gin-jwt-demo/pkg/config"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) Router(engine *gin.Engine) {
	engine.POST("/", middleware.JWTAuth(), uc.getUserInfo)
	engine.POST("/userLogin", uc.userLogin)
}

func (uc *UserController) getUserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"code":   0000,
		"msg":    "登录成功",
	})
}

// Login 登录
func (uc *UserController) userLogin(c *gin.Context) {
	var userReq model.LoginReq
	if e := c.Bind(&userReq); e != nil {
		c.JSON(200, gin.H{
			"code": 9999,
			"msg":  "数据格式错误",
		})
		return
	}

	if userReq.Mobile == "18954299351" {
		expiresTime := time.Now().Unix() + int64(10000000)
		mySigningKey := []byte(config.ConfigIni.Secret)
		// Create the Claims
		claims := &jwt.StandardClaims{
			Audience:  userReq.Mobile,    // 受众
			ExpiresAt: expiresTime,       // 失效时间
			Id:        "123",             // 编号
			IssuedAt:  time.Now().Unix(), // 签发时间
			Issuer:    "admin",           // 签发人
			NotBefore: time.Now().Unix(), // 生效时间
			Subject:   "login",           // 主题
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(mySigningKey)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "9999",
				"msg":    "登录失败",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "0000",
				"msg":    "登录成功",
				"token":  fmt.Sprintf("Bearer %v", ss),
			})
		}
		return
	}
	c.JSON(200, gin.H{
		"status": 9999,
		"msg":    "找不到用户",
	})
}
