/*
 * @Author: zhengzhuang
 * @Date: 2021-07-16 16:31:29
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-17 15:35:07
 * @Description: 中间件
 * @FilePath: /01-study/gin-demo/middleware/global_middleware.go
 */
package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 全局中间件
func GlobalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// fmt.Println("全局中间件开始执行了=======》》》》")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "全局中间件")
		// 执行函数
		c.Next()
		// 中间件执行完后续的一些事情
		// status := c.Writer.Status()
		// fmt.Println("全局中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}
