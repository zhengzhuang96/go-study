// Author: zhengzhuang
// Date: 2021-07-16 15:25:50
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-21 21:43:59
// Description: 渲染	ginRender
// FilePath: /01-study/gin-demo/controller/gin_render_controller.go
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

type GinRenderController struct{}

func (grc *GinRenderController) Router(engine *gin.Engine) {
	engine.GET("/someJson", grc.someJSON)
	engine.GET("/someStruct", grc.someStruct)
	engine.GET("/someXml", grc.someXML)
	engine.GET("/someYaml", grc.someYAML)
	engine.GET("/someProtoBuf", grc.someProtoBuf)
}

// 1.json
func (grc *GinRenderController) someJSON(c *gin.Context) {
	// fmt.Println("响应格式")
	c.JSON(200, gin.H{"message": "someJSON", "status": 200})
}

// 2.结构体响应格式
func (grc *GinRenderController) someStruct(c *gin.Context) {
	var msg struct {
		Name    string
		Message string
		Number  int
	}

	msg.Name = "名字a"
	msg.Message = "内容b"
	msg.Number = 333
	c.JSON(200, msg)
}

// 3.XML
func (grc *GinRenderController) someXML(c *gin.Context) {
	c.XML(200, gin.H{"name": "zhangsan"})
}

// 4.YAML
func (grc *GinRenderController) someYAML(c *gin.Context) {
	c.YAML(200, gin.H{"name": "zhangsan"})
}

// 5.protoBuf格式，谷歌开发的高效存储读取的工具
// 自定义传输格式
func (grc *GinRenderController) someProtoBuf(c *gin.Context) {
	reps := []int64{int64(1), int64(2)}
	// 定义数组
	label := "label"
	// 传protoBuf格式数据
	data := &protoexample.Test{
		Label: &label,
		Reps:  reps,
	}
	c.ProtoBuf(200, data)
}
