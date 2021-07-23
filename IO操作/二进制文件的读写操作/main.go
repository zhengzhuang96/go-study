// Author: zhengzhuang
// Date: 2021-07-23 15:57:20
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-23 15:57:20
// Description: 二进制文件的读写操作
// FilePath: /01-study/IO操作/二进制文件的读写操作/main.go
package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	// createFile()
	readFile()
}

// 创建并写入一个二进制文件
func createFile() {
	info := "golanggolanggolanggolanggolang"
	file, err := os.Create("./output.gob")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer file.Close()

	encode := gob.NewEncoder(file)
	err = encode.Encode(info)
	if err != nil {
		fmt.Println("编码错误", err.Error())
		return
	}
	fmt.Println("编码成功")
}

// 读取二进制文件
func readFile() {
	file, err := os.Open("./output.gob")
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return
	}
	defer file.Close()

	decode := gob.NewDecoder(file)
	info := ""
	err = decode.Decode(&info)
	if err != nil {
		fmt.Println("解码失败:", err.Error())
		return
	}
	fmt.Println(info)
}
