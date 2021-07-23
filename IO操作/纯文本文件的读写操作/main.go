// Author: zhengzhuang
// Date: 2021-07-23 15:41:36
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-23 15:41:37
// Description: 纯文本文件的读写操作
// FilePath: /01-study/IO操作/纯文本文件的读写操作/main.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// createFile()
	readFile()
}

// 创建并写如一个文件
func createFile() {
	filepath := "./output.txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("打开文件错误=%v \n", err)
		return
	}
	defer file.Close()
	// 写入内容
	str := "我是写入的新内容 \n"
	// 写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	// 写入3边
	for i := 0; i < 3; i++ {
		writer.WriteString(str)
	}
	// 因为 writer 是带缓存的，因此在调用 WriterString 方法时，内容是先写入缓存的
	// 所以要调用 flush 方法，将混存的数据真正写入到文件中
	writer.Flush()
}

// 读取文本内容
func readFile() {
	// 打开文件
	file, err := os.Open("./output.txt")
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	defer file.Close()
	// 创建一个 *Reader，是带缓存的
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 读到一个执行就结束
		if err == io.EOF {                  // io.EOF 表示文件的末尾
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束...")
}
