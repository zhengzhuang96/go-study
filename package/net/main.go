// Author: zhengzhuang
// Date: 2021-07-26 11:22:50
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-26 11:23:18
// Description: In User Settings Edit
// FilePath: /go-study/package/net/main.go

package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// dialHandle()
	// searchIp()
	TcpServer()
}

// Dial函数和服务器建立链接
func dialHandle() {
	conn, err := net.Dial("tcp", "baidu.com:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	// NewReader返回其缓冲区具有默认大小的新读取器。
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}

// 根据域名查找IP地址
func searchIp() {
	domain := "www.baidu.com"
	ip, err := net.ResolveIPAddr("ip", domain)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s的ip地址是:%s\n", domain, ip)
}

// 简单的TCP服务器
func TcpServer() {
	listen, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	for {
		con, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go TcpRead(con)
	}
}

func TcpRead(con net.Conn) {
	data := make([]byte, 1000)
	for {
		n, err := con.Read(data)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(data[0:n]))
	}
}
