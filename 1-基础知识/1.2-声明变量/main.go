// Author: zhengzhuang
// Date: 2021-07-29 14:39:06
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-29 14:39:06
// Description: In User Settings Edit
// FilePath: /go-study/1-基础知识/1.2-声明变量/main.go
package main

func main() {
	/*
	   一行声明一个变量
	   var <name> <type>
	   var 变量名 类型
	   其中 var 是关键字（固定不变），name 是变量名，type 是类型。
	*/
	var name string = "声明变量" // 变量声明后必须使用
	println(name)

	// 也可简写
	var name2 = "简写声明变量"
	println(name2)

	/*
		多个变量一起声明
		声明多个变量，除了可以按照上面写成多行之外，还可以写成下面这样
	*/
	var (
		name3  string
		age    int
		gender string
	)
	println(name3, age, gender)

	/*
		声明和初始化一个变量, 即：类型推导
		使用 := （推导声明写法或者短类型声明法：编译器会自动根据右值类型推断出左值的对应类型。），可以声明一个变量，并对其进行（显式）初始化。
		但这种方法有个限制就是，只能用于函数内部
	*/
	name4 := "类型推导" // 只能作用于函数内部
	println(name4)

	/*
		声明和初始化多个变量，
		这个方法，也经常用于变量的交换
	*/
	name5, age2 := "声明多个变量", 28
	println(name5, age2)

	// 变量的交换
	var a int = 100
	var b int = 200
	b, a = a, b
	println(a, b)

}
