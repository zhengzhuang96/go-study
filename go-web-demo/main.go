// Author: zhengzhuang
// Date: 2021-07-28 15:46:31
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-28 15:46:32
// Description: In User Settings Edit
// FilePath: /go-study/go-web-demo/main.go
package main

import (
	"fmt"
	"go-web-demo/cgo"
	"go-web-demo/routers"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		// Addr可选地指定服务器要侦听的TCP地址，
		// 以“主机:端口”的形式。如果为空，则使用“:http”（端口80）。
		// 服务名称在RFC 6335中定义并由IANA分配。
		// 有关地址格式的详细信息，请参阅net.Dial。
		Addr:    ":8080",
		Handler: cgo.Router,
		// 超时
		ReadTimeout: 5 * time.Second,
	}

	routers.SetupRouter(cgo.Router)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("程序完美运行...")

}
