// Author: zhengzhuang
// Date: 2021-07-24 14:42:00
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-24 14:42:01
// Description: In User Settings Edit
// FilePath: /01-study/im-system-demo/main.go
package main

func main() {
	server := NewServer("127.0.0.1", "8888")
	server.Start()
}
