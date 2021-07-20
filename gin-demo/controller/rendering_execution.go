/*
 * @Author: zhengzhuang
 * @Date: 2021-07-16 16:12:04
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-16 16:24:43
 * @Description: gin渲染执行
 * @FilePath: /gin-demo/app/rendering_execution.go
 */
package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 重定向，定一个重定向地址，访问会跳转目标地址
func Redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
}

// 异步
// goroutine机制可以方便地实现异步处理
// 另外，在启动新的goroutine时，不应该使用原始上下文，必须使用它的只读副本
func LongAsync(c *gin.Context) {
	// 这个不能直接使用c，而是搞一个副本
	copyContext := c
	// 异步处理
	go func() {
		// 等待3s后执行
		time.Sleep(3 * time.Second)
		log.Println("异步执行:" + copyContext.Request.URL.Path)
	}()
}

// 同步
func LongSync(c *gin.Context) {
	time.Sleep(3 * time.Second)
	log.Println("同步执行:" + c.Request.URL.Path)
}
