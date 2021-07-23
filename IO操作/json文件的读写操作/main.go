// Author: zhengzhuang
// Date: 2021-07-23 14:51:11
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-23 14:56:01
// Description: In User Settings Edit
// FilePath: /01-study/IO操作/json文件的读写操作/main.go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 创建一个json文件
type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

func main() {
	// createJson()
	readJson()
}

// 创建json
func createJson() {
	info := []Website{{
		"哈哈哈", "json文本", []string{"就送达阿斯顿执行", "我是文本"},
	}, {
		"golang", "文本呢吧", []string{"哈哈哈哈", "嫩嫩额"},
	}}

	// 创建文件
	filepath, err := os.Create("info.json")
	if err != nil {
		fmt.Println("文件创建失败,", err.Error())
		return
	}
	defer filepath.Close()

	// 创建json编码器, 并插入
	encode := json.NewEncoder(filepath)
	err = encode.Encode(info)
	if err != nil {
		fmt.Println("编码错误", err.Error())
		return
	}
	fmt.Println("编码成功")
}

// 读取json
func readJson() {
	fileStr, err := os.Open("./info.json")
	if err != nil {
		fmt.Println("文件打开失败:", err.Error())
		return
	}
	defer fileStr.Close()

	var info []Website
	// 创建json解码器
	decode := json.NewDecoder(fileStr)
	err = decode.Decode(&info)
	if err != nil {
		fmt.Println("解码失败", err.Error())
		return
	}
	fmt.Println("解码成功")
	fmt.Println(info)
}
