/*
 * @Author: zhengzhuang
 * @Date: 2021-07-14 11:54:46
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-14 14:03:27
 * @Description: byte和rune类型
 * @FilePath: /day01/07byte和rune类型/main.go
 */
package main

import (
	"fmt"
	"unicode"
)

// 组成每个字符串的元素叫做“字符”，可以通过遍历或者单个获取字符串元素获得字符。 字符用单引号（’）包裹起来，如：
var a = '中'
var b = 'x'

// Go 语言的字符有以下两种：
// uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
// rune类型，代表一个 UTF-8字符。
// 当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32。

func main() {
	// traversalString()
	// changeString()
	test()
}

// 遍历字符串
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { // rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

// 修改字符串
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}

// 分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Prinf()搭配%T分别打印出上述变量的值和类型
func test() {
	var (
		a = 10
		b = 1.121
		c = false
		d = "hello"
	)
	fmt.Printf("%T %v \n", a, a)
	fmt.Printf("%T %v\n", b, b)
	fmt.Printf("%T %v\n", c, c)
	fmt.Printf("%T %v\n", d, d)

	e := "hello沙河小王子啊"

	fmt.Println("===============")

	f := 0
	for _, char := range e {
		if unicode.Is(unicode.Han, char) {
			f++
		}
	}
	fmt.Println(f)
}
