// Author: zhengzhuang
// Date: 2021-07-23 16:06:12
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-23 16:06:13
// Description: zip归档文件的读写操作
// FilePath: /01-study/IO操作/zip归档文件的读写操作/main.go
package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// createZip()
	readZip()
}

// 创建zip归档文件
func createZip() {
	// 创建一个缓冲区用来保存压缩文件内容
	buf := new(bytes.Buffer)

	// 创建一个压缩文档
	w := zip.NewWriter(buf)

	// 将文件加入压缩文档
	var files = []struct {
		Name, Body string
	}{
		{"Golang.txt", "我是写入的文件内容1"},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			fmt.Println(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			fmt.Println(err)
		}
	}

	// 关闭压缩文档
	err := w.Close()
	if err != nil {
		fmt.Println(err)
	}

	// 将压缩文档内容写入文件
	f, err := os.OpenFile("file.zip", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}
	buf.WriteTo(f)
}

// 读取 zip 归档文件
func readZip() {
	// 打开一个zip格式文件
	r, err := zip.OpenReader("file.zip")
	if err != nil {
		fmt.Println("压缩文件打开失败", err)
		return
	}
	defer r.Close()

	// 迭代压缩文件中的文件，打印出文件中的内容
	for _, v := range r.File {
		fmt.Printf("文件名: %s \n", v.Name)
		rc, err := v.Open()
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = io.CopyN(os.Stdout, rc, int64(v.UncompressedSize64))
		if err != nil {
			fmt.Printf(err.Error())
		}
		rc.Close()
	}
}
