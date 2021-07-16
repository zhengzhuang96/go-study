/*
 * @Author: zhengzhuang
 * @Date: 2021-07-15 16:59:27
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-15 17:10:22
 * @Description: 匿名函数
 * @FilePath: /函数/匿名函数/main.go
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a) // 返回 a 的平方根
	}

	fmt.Println(getSqrt(4))

	// ------- 函数变量 ------
	println("------- 函数变量 ------")
	fn := func() { println("Hello, World!") }
	fn()

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

	// --------- channel of function -----------
	fc := make(chan func() string, 2)
	fc <- func() string { return "Hello, World!" }
	println((<-fc)())
}
