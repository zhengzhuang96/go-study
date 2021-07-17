/*
 * @Author: zhengzhuang
 * @Date: 2021-07-16 10:40:02
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-17 15:57:17
 * @Description: In User Settings Edit
 * @FilePath: /01-study/gin-demo/routers/routers.go
 */
package routers

import (
	"ginDemo/app"
	"ginDemo/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 注册中间件
	r.Use(middleware.GlobalMiddleware())
	// form表单的测试html
	r.GET("/", app.FormHtml)
	// 通过Context的Param方法来获取API参数
	r.GET("/apiParameters/:name/*action", app.ApiParameters)
	// URL参数	来获取?后的值
	r.GET("/urlParameters", app.UrlParameters)
	// 表单参数可以通过PostForm()方法获取
	r.POST("/formParameters", app.FormParameters)
	// 上传单个文件
	r.POST("/uploadFile", app.UploadFile)
	// 上传多个文件
	r.POST("/uploadMultipleFiles", app.UploadMultipleFiles)
	// json数据解析与绑定
	r.POST("/jsonDataHandle", app.JsonDataHandle)
	// form表单数据解析与绑定
	r.POST("/formDataHandle", app.FormDataHandle)
	// url数据解析与绑定
	r.GET("/urlDataHandle/:user/:password", app.UrlDataHandle)
	// json响应格式
	r.GET("/someJson", app.SomeJSON)
	// 结构体响应格式
	r.GET("/someStruct", app.SomeStruct)
	// XML响应格式
	r.GET("/someXml", app.SomeXML)
	// YAML响应格式
	r.GET("/someYaml", app.SomeYAML)
	// protoBuf响应格式
	r.GET("/someProtoBuf", app.SomeProtoBuf)
	// HTML模板渲染
	r.GET("/htmlTemplate", app.HtmlTemplate)
	// 路由重定向
	r.GET("/redirect", app.Redirect)
	// 异步
	r.GET("/longAsync", middleware.LocalMiddleware(), app.LongAsync)
	// 同步
	r.GET("/longSync", app.LongSync)
	// cookie练习
	r.GET("/cookie", app.Cookie)
	// 获取并设置cookie
	r.GET("/loginGetCookie", app.LoginSetCookie)
	// 需要cookie的页面
	r.GET("/homeCookie", middleware.AuthMiddleWare(), app.HomeCookie)
	// 往数据库添加数据
	r.GET("/addData", app.AddData)
	// 删除数据库数据
	r.GET("/deleteData", app.DeleteData)
	// 更新数据库数据
	r.GET("/updateData", app.UpdateData)
	// 超找数据库
	r.GET("/selectData", app.SelectData)
	// 存数据到redis
	r.GET("/connectRedis", app.ConnectRedis)
	// 从redis里面取数据
	r.GET("/getRedis", app.GetRedis)
	// 存多个数据到redis
	r.GET("/batchRedis", app.BatchRedis)
	// 取多个redis中的值
	r.GET("/getBatchRedis", app.GetBatchRedis)
	// 在redis中存储一个有过期时间的值
	r.GET("/expireTimeRedis", app.ExpireTimeRedis)
	return r
}
