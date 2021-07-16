/*
 * @Author: zhengzhuang
 * @Date: 2021-07-16 15:51:33
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-16 16:10:08
 * @Description: In User Settings Edit
 * @FilePath: /gin-demo/app/html_template.go
 */
package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//form表单的测试
func FormHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", gin.H{"title": "我是测试", "ce": "123456"})
}

// HTML模板渲染
func HtmlTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "ce": "123456"})
}
