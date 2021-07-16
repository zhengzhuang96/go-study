/*
 * @Author: zhengzhuang
 * @Date: 2021-07-13 22:37:04
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-14 14:06:34
 * @Description: 声明变量
 * @FilePath: /day01/02var(变量声明和赋值)/main.go
 */
package main

import "fmt"

// 推荐使用驼峰式命名

// 声明变量, 单独声明
// var 变量名 类型 = 表达式
// var name string
// var age int
// var isOk bool

// 批量声明
var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)

// 全局变量m
var m = 100

func main() {

	name = "理想"
	age = 16
	isOk = true
	fmt.Print(age)              // 在终端输出要打印的内容
	fmt.Printf("isOk:%s", name) // %s: 占位符 使用name这个变量的值去替换占位符
	fmt.Println(age)            // 打印完指定的内容之后会在后面加一个换行符

	// 变量声明
	var s0 string = "asd"
	fmt.Println(s0)

	// 类型推导（根据值判断变量类型）
	var s2 = "20"
	fmt.Println(s2)

	// 简短变量声明（声明变量同时赋值，根据存储值判断类型）
	s1 := "哈哈哈"
	fmt.Println(s1)
}
