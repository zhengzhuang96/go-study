// Author: zhengzhuang
// Date: 2021-07-16 17:32:41
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-21 21:45:25
// Description: cookie 验证
// FilePath: /01-study/gin-demo/middleware/auth_middleware.go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并效验
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
	}
}
