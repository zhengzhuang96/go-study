/*
 * @Author: zhengzhuang
 * @Date: 2021-07-15 16:59:27
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-19 09:30:08
 * @Description: 匿名函数
 * @FilePath: /01-study/函数/匿名函数/main.go
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	// GeneralAnonymity()
	// FunctionSet()
	// FieldFunction()
	// ChannelFunction()
	AnonymousCallback()
}

// 普通匿名函数
func GeneralAnonymity() {
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a) // 返回 a 的平方根
	}

	fmt.Println(getSqrt(4))

	// ------- 函数变量 ------
	println("------- 函数变量 ------")
	fn := func() { println("Hello, World!") }
	fn()
}

// 函数集合
func FunctionSet() {
	// -------- 函数集合 -----------
	println("-------- 函数集合 -----------")
	fns := [](func(x int) int){
		func(x int) int {
			return x + 1
		},
		func(x int) int {
			return x + 2
		},
	}
	println(fns[0](100))
}

// 字段函数
func FieldFunction() {
	// ------- 作为字段的函数 --------
	println("------- 作为字段的函数 --------")
	d := struct {
		fn func() string
	}{
		fn: func() string {
			return "Hello, World!"
		},
	}
	println(d.fn())
}

// channel of function
func ChannelFunction() {
	fc := make(chan func() string, 2)
	fc <- func() string { return "Hello, World!" }
	println((<-fc)())
}

// 匿名回调函数
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}
func AnonymousCallback() {
	l := []int{1, 2, 3}
	visit(l, func(i int) {
		fmt.Println(i)
	})
}
