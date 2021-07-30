// Author: zhengzhuang
// Date: 2021-07-27 16:07:03
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-27 16:07:36
// Description: In User Settings Edit
// FilePath: /go-study/gin/gin-jwt-demo/middleware/jwt.go
package middleware

import (
	"gin-jwt-demo/pkg/config"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.PostForm("token")
		if auth == "" {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 9999,
				"msg":    "无法认证，重新登录",
			})
			return
		}
		// auth = strings.Fields(auth)[1]
		// 校验token
		auth = string(auth)[7:]
		_, err := parseToken(auth)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 9999,
				"msg":    "token 过期",
			})
		} else {
			println("token 正确")
		}
		c.Next()
	}
}

func parseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(config.ConfigIni.Secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
