// Author: zhengzhuang
// Date: 2021-07-23 16:47:21
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-23 16:47:22
// Description: buffer读取文件
// FilePath: /01-study/IO操作/buffer读取文件/main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// writeFile()
	readFile()
}

// 写入 bufio 文件
func writeFile() {
	name := "demo.txt"
	context := "我是写入demo里的文件数据。。啦啦啦"

	fileObj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return
	}
	defer fileObj.Close()
	writeObj := bufio.NewWriterSize(fileObj, 4096)

	// 使用 Write 方法，需要使用 Write 对象的 Flush 方法将 buffer 中的数据刷到磁盘
	buf := []byte(context)
	if _, err := writeObj.Write(buf); err == nil {
		if err := writeObj.Flush(); err != nil {
			panic(err)
		}
		fmt.Println("数据写入成功")
	}
}

// 使用 bufio 包读取文件
func readFile() {
	fileObj, err := os.Open("demo.txt")
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return
	}
	defer fileObj.Close()

	//一个文件对象本身是实现了io.Reader的 使用bufio.NewReader去初始化一个Reader对象，存在buffer中的，读取一次就会被清空
	reader := bufio.NewReader(fileObj)
	buf := make([]byte, 1024)
	// 读取Reader对象中的内容到[]byte类型的buf中
	info, err := reader.Read(buf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("读取的字节数:", strconv.Itoa(info))
	// 这里的buf是一个[]byte, 因此如果需要只输出内容，仍然需要将文件内容的换行符替换掉
	fmt.Println("读取的文件内容:", string(buf))
}
