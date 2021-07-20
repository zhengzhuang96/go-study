/*
 * @Author: zhengzhuang
 * @Date: 2021-07-16 10:37:19
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-20 17:59:19
 * @Description: 主程序
 * @FilePath: /01-study/gin-demo/main.go
 */
package main

import (
	"fmt"
	"gin-demo/database"
	"gin-demo/middleware"
	"io"
	"os"

	"github.com/gin-gonic/gin"

	"gin-demo/routers"
	"gin-demo/tool"

	_ "github.com/go-sql-driver/mysql"
)

// 在main之前运行
func init() {
	database.Conn()
}

// 主程序
func main() {
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}

	_, err = tool.OrmEngine(cfg)
	if err != nil {
		panic(err.Error())
	}

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

	// 设置全局跨域访问
	r.Use(middleware.Cors())

	if err := r.Run(cfg.AppHost + ":" + cfg.AppPort); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
