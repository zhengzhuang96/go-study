/*
 * @Author: zhengzhuang
 * @Date: 2021-07-14 08:55:42
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-14 10:00:59
 * @Description: 常量
 * @FilePath: /day01/03const(常量)/main.go
 */
package main

import "fmt"

// 常量
// 定义了常量之后不能修改
// 在程序运行期间不会改变的量
const pi = 3.1415926

// 批量声明常量
const (
	statusOk = 200
	notFound = 404
)

// 批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一致
const (
	n1 = 100 // 0
	n2       // 1
	n3       // 2
)

// iota
const (
	a1 = iota // 0
	a2 = iota // 1
	a3        // 2
)

const (
	b1 = iota // 0
	b2        // 1
	_
	b3 // 3
)

// 插队
const (
	c1 = iota // 0
	c2 = 100  // 100
	c3 = iota // 2
	c4
)

// 多个常量声明在一行
// 在const出现的时候ioat被置为0，每加一行+1，如果在一行不变
const (
	d1, d2 = iota + 1, iota + 2 // d1:1 d2:2
	d3, d4 = iota + 1, iota + 2 // d3:2 d4:3
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	// 常量赋值之后不允许修改
	// pi = 123123	// 不被允许

	// fmt.Println("n1", n1)
	// fmt.Println("n2", n2)
	// fmt.Println("n3", n3)

	// fmt.Println("a1", a1)
	// fmt.Println("a2", a2)
	// fmt.Println("a3", a3)

	// fmt.Println("b1", b1)
	// fmt.Println("b2", b2)
	// fmt.Println("b3", b3)

	// fmt.Println("c1", c1)
	// fmt.Println("c2", c2)
	// fmt.Println("c3", c3)
	// fmt.Println("c4", c4)

	fmt.Println("d1", d1)
	fmt.Println("d2", d2)
	fmt.Println("d3", d3)
	fmt.Println("d4", d4)
}
