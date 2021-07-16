/*
 * @Author: zhengzhuang
 * @Date: 2021-07-16 13:58:58
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-16 15:04:56
 * @Description: 数据解析与绑定
 * @FilePath: /gin-demo/app/dataParsingBinding.go
 */
package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义接收数据的结构体
type Login struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

// json数据解析与绑定
// curl http://127.0.0.1:8080/jsonData -H 'content-type:application/json' -d "{\"User\": \"root\", \"password\": \"admin\"}" -X POST
func JsonDataHandle(c *gin.Context) {
	// 1.声明接收的变量
	var json Login
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&json); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 判断用户名密码是否正确
	if json.User != "root" || json.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

// form表单数据解析与绑定
func FormDataHandle(c *gin.Context) {
	// 声明接收的变量
	var form Login
	// Bind()默认解析并绑定form格式
	// 根据请求头中content-type自行推断
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if form.User != "root" || form.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"status": "200"})
}

// url数据解析与绑定
func UrlDataHandle(c *gin.Context) {
	// 声明接收的变量
	var login Login
	// Bind()默认解析并绑定form格式
	// 根据请求头中content-type自动推断
	if err := c.ShouldBindUri(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if login.User != "root" || login.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}
