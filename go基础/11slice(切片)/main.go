/*
 * @Author: zhengzhuang
 * @Date: 2021-07-14 16:29:16
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-14 16:33:58
 * @Description: 切片	main
 * @FilePath: /day01/11slice(切片)/main.go
 */
package main

import "fmt"

// 切片
// 切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
// 切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
// 声明切片类型的基本语法如下：
// var name []T
// var name:变量名	T切片中的元素类型

// 切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。

func main() {
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	// var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)        // []
	fmt.Println(b)        // []
	fmt.Println(c)        // [false true]
	fmt.Println(a == nil) // true
	fmt.Println(b == nil) // false
	fmt.Println(c == nil) // false
	// fmt.Println(c == d)   // 切片是引用类型，不支持直接比较，只能和nil比较

	g := [5]int{1, 2, 3, 4, 5}
	s := g[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))

	fmt.Println(g[2:]) // 等同于 g[2:len(g)]
	fmt.Println(g[:3]) // 等同于 g[0:3]
	fmt.Println(g[:])  // 等同于 g[0:len(g)]
}
