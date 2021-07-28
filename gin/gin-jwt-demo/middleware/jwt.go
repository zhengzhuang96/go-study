// Author: zhengzhuang
// Date: 2021-07-27 16:07:03
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-27 16:07:36
// Description: In User Settings Edit
// FilePath: /go-study/gin/gin-jwt-demo/middleware/jwt.go
package middleware

import (
	"fmt"
	"gin-jwt-demo/tool"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "" {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		fmt.Println("get token:", token)

		j := tool.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == tool.TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理，并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}
