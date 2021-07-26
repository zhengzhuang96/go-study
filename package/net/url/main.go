// Author: zhengzhuang
// Date: 2021-07-26 14:54:04
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-26 14:54:21
// Description: net/url url包
// FilePath: /go-study/package/net/url/main.go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	//Parse将rawurl解析为URL结构。
	//rawurl可以是相对的（路径，没有主机）或绝对的
	//（从一个方案开始）。正在尝试分析主机名和路径
	//没有方案是无效的，但不一定返回
	u, err := url.Parse("http://www.baidu.com/search?q=dotnet")
	if err != nil {
		fmt.Println(err)
	}
	u.Scheme = "https"    // http替换为https
	u.Host = "google.com" // baidu.com替换为google.com
	q := u.Query()
	fmt.Println(q)
	q.Set("q", "golang")
	u.RawQuery = q.Encode() // dotnet替换为golang
	fmt.Println(u)
}
