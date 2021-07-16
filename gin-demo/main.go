/*
 * @Author: zhengzhuang
 * @Date: 2021-07-16 10:37:19
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-16 18:04:17
 * @Description: In User Settings Edit
 * @FilePath: /gin-demo/main.go
 */
package main

import (
	"fmt"
	"ginDemo/routers"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 日志文件
	gin.DisableConsoleColor()
	// 创建日志文件并录入文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// 这一行可以将日志显示到控制台
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 路由加载
	r := routers.SetupRouter()
	// LoadHTMLGlob()方法可以加载模板文件
	r.LoadHTMLGlob("template/*")
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
