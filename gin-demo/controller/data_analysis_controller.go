/*
 * @Author: zhengzhuang
 * @Date: 2021-07-16 13:58:58
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-20 21:10:33
 * @Description: 数据解析与绑定
 * @FilePath: /01-study/gin-demo/controller/data_analysis_controller.go
 */
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DataAnalysisController struct{}

// 定义接收数据的结构体
type Login struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func (dataAc *DataAnalysisController) Router(engine *gin.Engine) {
	engine.POST("/jsonData", dataAc.jsonDataHandle)
	engine.POST("/formDataHandle", dataAc.formDataHandle)
	engine.GET("/urlDataHandle/:user/:password", dataAc.urlDataHandle)
}

// json数据解析与绑定
// curl http://127.0.0.1:8080/jsonData -H 'content-type:application/json' -d "{\"User\": \"root\", \"password\": \"admin\"}" -X POST
func (dataAc *DataAnalysisController) jsonDataHandle(c *gin.Context) {
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
func (dataAc *DataAnalysisController) formDataHandle(c *gin.Context) {
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
func (dataAc *DataAnalysisController) urlDataHandle(c *gin.Context) {
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
