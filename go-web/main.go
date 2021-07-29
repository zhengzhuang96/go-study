// Author: zhengzhuang
// Date: 2021-07-29 08:33:14
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-29 08:33:15
// Description: In User Settings Edit
// FilePath: /go-study/go-web/main.go
package main

import (
	"go-web/cgo"
	"net/http"
)

func main() {
	r := cgo.New()

	r.GET("/", func(c *cgo.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	// r.POST("/hello", func(c *cgo.Context) {
	// 	fmt.Println("HELLO")
	// })

	r.Run(":8080")
}
