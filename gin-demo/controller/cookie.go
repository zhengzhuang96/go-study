/*
 * @Author: zhengzhuang
 * @Date: 2021-07-16 17:02:57
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-19 17:39:17
 * @Description: Cookie的使用
 * @FilePath: /01-study/gin-demo/controller/cookie.go
 */
package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Cookie
// HTTP是无状态协议，服务器不能记录浏览器的访问状态，也就是说服务器不能区分两次请求是否由同一个客户端发出
// Cookie就是解决HTTP协议无状态的方案之一，中文是小甜饼的意思
// Cookie实际上就是服务器保存在浏览器上的一段信息。浏览器有了Cookie之后，每次向服务器发送请求时都会同时将该信息发送给服务器，服务器收到请求后，就可以根据该信息处理请求
// Cookie由服务器创建，并发送给浏览器，最终由浏览器保存
// Cookie的用途
// 测试服务端发送cookie给客户端，客户端请求时携带cookie
// Cookie的缺点
// 不安全，明文
// 增加带宽消耗
// 可以被禁用
// cookie有上限

// 设置一个cookie
func Cookie(c *gin.Context) {
	// 获取客户端是否携带cookie
	cookie, err := c.Cookie("key_cookie")
	if err != nil {
		cookie = "NotSet"
		// 给客户端设置cookie
		// maxAge int, 单位为秒
		// path, cookie所在目录
		// domain string, 域名
		// secure 是否智能通过https访问
		// httpOnly bool 是否允许别人通过js获取自己的cookie
		c.SetCookie("key_cookie", "value_cookie", 60, "/", "localhost", false, true)
	}
	fmt.Printf("cookie的值是：%s\n", cookie)
}

// login登录设置cookie
func LoginSetCookie(c *gin.Context) {
	// 设置cookie
	c.SetCookie("abc", "123", 60, "/", "localhost", false, true)
	// 返回信息
	c.String(200, "Login success!!!")
}

// 访问需要cookie的页面
func HomeCookie(c *gin.Context) {
	c.JSON(200, gin.H{"data": "Home"})
}
