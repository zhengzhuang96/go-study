/*
 * @Author: zhengzhuang
 * @Date: 2021-07-14 10:42:37
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-14 14:08:22
 * @Description: int
 * @FilePath: /day01/04int/main.go
 */
package main

import "fmt"

// 整型
// 类型		 描述
// uint8	无符号 8位整型 	(0 到 255)
// uint16	无符号 16位整型 (0 到 65535)
// uint32	无符号 32位整型 (0 到 4294967295)
// uint64	无符号 64位整型 (0 到 18446744073709551615)
// int8		有符号 8位整型 	(-128 到 127)
// int16	有符号 16位整型 (-32768 到 32767)
// int32	有符号 32位整型 (-2147483648 到 2147483647)
// int64	有符号 64位整型 (-9223372036854775808 到 9223372036854775807)

// 特殊整型
// 类型			 描述
// uint			32位操作系统上就是uint32，64位操作系统上就是uint64
// int			32位操作系统上就是int32，64位操作系统上就是int64
// uintptr	无符号整型，用于存放一个指针

// 整型
func main() {

	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010占位符%b表示二进制

	// 八进制 以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制	以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF
}
