// Author: zhengzhuang
// Date: 2021-07-23 15:23:21
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-23 15:23:21
// Description: xml文件的读写操作
// FilePath: /01-study/IO操作/xml文件的读写操作/main.go
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Cource []string
}

func main() {
	// createXml()
	readXml()
}

// 创建并写入xml文件
func createXml() {
	// 实例化对象
	info := Website{"哈哈哈哈", "嗯哼", []string{"哦哦哦", "鞍钢那个"}}

	f, err := os.Create("./info.xml")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer f.Close()

	// 序列化到文件中
	encode := xml.NewEncoder(f)
	err = encode.Encode(info)
	if err != nil {
		fmt.Println("编码错误:", err.Error())
		return
	}
	fmt.Println("编码成功")
}

// 读xml文件
func readXml() {
	// 打开xml文件
	file, err := os.Open("./info.xml")
	if err != nil {
		fmt.Println("文件读取失败:", err.Error())
		return
	}
	defer file.Close()

	info := Website{}
	// 创建 xml 解码器
	decode := xml.NewDecoder(file)
	err = decode.Decode(&info)
	if err != nil {
		fmt.Println("解码失败:", err.Error())
		return
	}
	fmt.Println("解码成功")
	fmt.Println(info)
}
