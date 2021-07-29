// Author: zhengzhuang
// Date: 2021-07-29 14:51:33
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-29 14:51:33
// Description: In User Settings Edit
// FilePath: /go-study/1-基础知识/1.4-byte、rune与字符串/main.go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/*
		byte，占用1个节字，就 8 个比特位（2^8 = 256，因此 byte 的表示范围 0->255），所以它和 uint8 类型本质上没有区别，它表示的是 ACSII 表中的一个字符。
	*/
	var a byte = 65
	// 8进制写法: var a byte = '\101'     其中 \ 是固定前缀
	// 16进制写法: var a byte = '\x41'    其中 \x 是固定前缀
	var b uint8 = 66
	fmt.Printf("a 的值: %c \nb 的值: %c\n", a, b)
	// 或者使用 string 函数
	// fmt.Println("a 的值: ", string(a)," \nb 的值: ", string(b))

	// 在 ASCII 表中，由于字母 A 的ASCII 的编号为 65 ，字母 B 的ASCII 编号为 66，所以上面的代码也可以写成这样
	var c byte = 'A'
	var d uint8 = 'B'
	fmt.Printf("c 的值: %c \nd 的值: %c\n", c, d)

	/*
		rune，占用4个字节，共32位比特位，所以它和 uint32 本质上也没有区别。它表示的是一个 Unicode字符（Unicode是一个可以表示世界范围内的绝大部分字符的编码规范）。
	*/
	var e byte = 'A'
	var f rune = 'B'
	fmt.Printf("e 占用 %d 个字节数\nf 占用 %d 个字节数\n", unsafe.Sizeof(e), unsafe.Sizeof(f))

	/*
		byte 和 rune 都是字符类型，若多个字符放在一起，就组成了字符串
	*/
	var mystr01 string = "hello"
	var mystr02 [5]byte = [5]byte{104, 101, 108, 108, 111}
	fmt.Printf("mystr01: %s\n", mystr01)
	fmt.Printf("mystr02: %s\n", mystr02)

	// Go 语言的 string 是用 uft-8 进行编码的，英文字母占用一个字节，而中文字母占用 3个字节

	var country string = "hello,中国"
	fmt.Println(len(country)) // 12, "hello," -> 6个字符   "中国" -> 6个字符，汉字每个占3个字符
}
