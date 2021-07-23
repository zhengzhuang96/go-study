// Author: zhengzhuang
// Date: 2021-07-23 17:33:58
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-23 17:33:59
// Description: 文件锁操作
// FilePath: /01-study/IO操作/文件锁操作/main.go
package main

import "fmt"

func main() {
	go fmt.Println("哈哈哈1")
	go fmt.Println("哈哈哈2")
}
