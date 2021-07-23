// Author: zhengzhuang
// Date: 2021-07-23 16:28:54
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-23 16:28:54
// Description: tar归档文件的读写操作
// FilePath: /01-study/IO操作/tar归档文件的读写操作/main.go
package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

func main() {
	// createFile()
	decompressionTar()
}

// 创建tar归档文件
func createFile() {
	// 创建一个tar文件
	f, err := os.Create("./output.tar")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	tw := tar.NewWriter(f)
	defer tw.Close()

	// 获取文件相关信息
	fileinfo, err := os.Stat("./txt/file.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	hdr, err := tar.FileInfoHeader(fileinfo, "")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 写入头文件信息
	err = tw.WriteHeader(hdr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	f1, err := os.Open("./txt/file.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	m, err := io.Copy(tw, f1) // 将file.txt 文件中的信息写入压缩包中
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(m)
}

// 解压tar归档文件
func decompressionTar() {
	f, err := os.Open("output.tar")
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return
	}
	defer f.Close()
	r := tar.NewReader(f)
	for hdr, err := r.Next(); err != io.EOF; hdr, err = r.Next() {
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fileinfo := hdr.FileInfo()
		fmt.Println(fileinfo.Name())
		//
		f, err := os.Create(fileinfo.Name())
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		_, err = io.Copy(f, r)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
