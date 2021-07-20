/*
 * @Author: zhengzhuang
 * @Date: 2021-07-16 10:40:02
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-19 17:41:38
 * @Description: In User Settings Edit
 * @FilePath: /01-study/gin-demo/routers/routers.go
 */
package routers

import (
	"gin-demo/controller"
	"gin-demo/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 注册中间件
	r.Use(middleware.GlobalMiddleware())
	// form表单的测试html
	r.GET("/", controller.FormHtml)
	// 通过Context的Param方法来获取API参数
	r.GET("/apiParameters/:name/*action", controller.ApiParameters)
	// URL参数	来获取?后的值
	r.GET("/urlParameters", controller.UrlParameters)
	// 表单参数可以通过PostForm()方法获取
	r.POST("/formParameters", controller.FormParameters)
	// 上传单个文件
	r.POST("/uploadFile", controller.UploadFile)
	// 上传多个文件
	r.POST("/uploadMultipleFiles", controller.UploadMultipleFiles)
	// json数据解析与绑定
	r.POST("/jsonDataHandle", controller.JsonDataHandle)
	// form表单数据解析与绑定
	r.POST("/formDataHandle", controller.FormDataHandle)
	// url数据解析与绑定
	r.GET("/urlDataHandle/:user/:password", controller.UrlDataHandle)
	// json响应格式
	r.GET("/someJson", controller.SomeJSON)
	// 结构体响应格式
	r.GET("/someStruct", controller.SomeStruct)
	// XML响应格式
	r.GET("/someXml", controller.SomeXML)
	// YAML响应格式
	r.GET("/someYaml", controller.SomeYAML)
	// protoBuf响应格式
	r.GET("/someProtoBuf", controller.SomeProtoBuf)
	// HTML模板渲染
	r.GET("/htmlTemplate", controller.HtmlTemplate)
	// 路由重定向
	r.GET("/redirect", controller.Redirect)
	// 异步
	r.GET("/longAsync", middleware.LocalMiddleware(), controller.LongAsync)
	// 同步
	r.GET("/longSync", controller.LongSync)
	// cookie练习
	r.GET("/cookie", controller.Cookie)
	// 获取并设置cookie
	r.GET("/loginGetCookie", controller.LoginSetCookie)
	// 需要cookie的页面
	r.GET("/homeCookie", middleware.AuthMiddleWare(), controller.HomeCookie)
	// 往数据库添加数据
	r.GET("/addData", controller.AddData)
	// 删除数据库数据
	r.GET("/deleteData", controller.DeleteData)
	// 更新数据库数据
	r.GET("/updateData", controller.UpdateData)
	// 超找数据库
	r.GET("/selectData", controller.SelectData)
	// 存数据到redis
	r.GET("/connectRedis", controller.ConnectRedis)
	// 从redis里面取数据
	r.GET("/getRedis", controller.GetRedis)
	// 存多个数据到redis
	r.GET("/batchRedis", controller.BatchRedis)
	// 取多个redis中的值
	r.GET("/getBatchRedis", controller.GetBatchRedis)
	// 在redis中存储一个有过期时间的值
	r.GET("/expireTimeRedis", controller.ExpireTimeRedis)
	return r
}
