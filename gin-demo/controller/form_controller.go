/*
 * @Author: zhengzhuang
 * @Date: 2021-07-16 11:33:18
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-20 21:20:32
 * @Description: 表单参数
 * @FilePath: /01-study/gin-demo/controller/form_controller.go
 */
package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 基础表单
// 表单传输为post请求，http常见的传输格式为四种：
// application/json
// application/x-www-form-urlencoded
// application/xml
// multipart/form-data

type FormController struct{}

func (form *FormController) Router(engine *gin.Engine) {
	engine.POST("/formParameters", form.formParameters)
	engine.POST("/uploadFile", form.uploadFile)
	engine.POST("/uploadMultipleFiles", form.uploadMultipleFiles)
}

// 表单参数可以通过PostForm()方法获取，该方法默认解析的是x-www-form-urlencoded或from-data格式的参数
func (form *FormController) formParameters(c *gin.Context) {
	fmt.Println("表单参数")
	types := c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.String(http.StatusOK, fmt.Sprintf("username: %s, password: %s, type: %s", username, password, types))
}

// 上传单个文件
func (form *FormController) uploadFile(c *gin.Context) {
	fmt.Println("上传单个文件")
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "文件上传失败")
		return
	}
	fmt.Println(file.Filename)
	c.SaveUploadedFile(file, "assets/"+file.Filename)
	c.String(http.StatusOK, "文件上传成功")
}

// 上传多个文件
func (form *FormController) uploadMultipleFiles(c *gin.Context) {
	// fmt.Println("上传多个文件")
	// form, err := c.MultipartForm()
	// if err != nil {
	// 	c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
	// }
	// // 获取所有图片
	// files := form.File["files"]
	// // 遍历所有图片
	// for _, file := range files {
	// 	// 逐个存
	// 	if err := c.SaveUploadedFile(file, "assets/"+file.Filename); err != nil {
	// 		c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
	// 		return
	// 	}
	// }
	// c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
}
