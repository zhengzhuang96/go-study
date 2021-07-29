// Author: zhengzhuang
// Date: 2021-07-29 14:22:00
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-29 14:22:01
// Description: In User Settings Edit
// FilePath: /go-study/01-基础知识/01-01:整数和浮点型/main.go
package main

import "fmt"

func main() {
	var num int = 10
	println(num)

	var num01 int = 0b1100
	var num02 int = 0o14
	var num03 int = 0xC

	fmt.Printf("2进制数 %b 表示的是: %d \n", num01, num01)
	fmt.Printf("8进制数 %o 表示的是: %d \n", num02, num02)
	fmt.Printf("16进制数 %X 表示的是: %d \n", num03, num03)

	var myfloat float32 = 10000018
	fmt.Println("myfloat: ", myfloat)
	fmt.Println("myfloat: ", myfloat+1)

	var myfloat01 float32 = 100000182
	var myfloat02 float32 = 100000187
	fmt.Println("myfloat: ", myfloat01)
	fmt.Println("myfloat: ", myfloat01+5)
	fmt.Println(myfloat02 == myfloat01+5)
}
