/*
 * @Author: zhengzhuang
 * @Date: 2021-07-16 15:51:33
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-20 10:42:45
 * @Description: In User Settings Edit
 * @FilePath: /01-study/gin-demo/controller/HtmlTemplateController.go
 */
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HtmlTemplateController struct{}

func (html *HtmlTemplateController) Router(engine *gin.Engine) {
	engine.GET("/", html.formHtml)
	engine.GET("/htmlTemplate", html.htmlTemplate)
}

//form表单的测试
func (html *HtmlTemplateController) formHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", gin.H{"title": "我是测试", "ce": "123456"})
}

// HTML模板渲染
func (html *HtmlTemplateController) htmlTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "ce": "123456"})
}
