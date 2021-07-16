/*
 * @Author: zhengzhuang
 * @Date: 2021-07-14 14:19:38
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-14 15:05:57
 * @Description: for循环
 * @FilePath: /day01/09for/main.go
 */
package main

import "fmt"

// for 初始语句;条件表达式;结束语句{
// 	循环体语句
// }

// for循环可以通过break、goto、return、panic语句强制退出循环。

// for循环
func main() {
	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("======变种1========")
	// 变种1	for 循环向上查找i定义的数值
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("======变种2========")
	// 变种2
	var y = 5
	for y < 10 {
		fmt.Println(y)
		y++
	}

	// fmt.Println("=======无限循环=======")
	// 无限循环
	// for {
	// 	fmt.Println("11111")
	// }

	fmt.Println("=======for range(键值循环)=======")
	// for range(键值循环)
	// 	Go语言中可以使用for range遍历数组、切片、字符串、map 及通道（channel）。 通过for range遍历的返回值有以下规律：
	// 1, 数组、切片、字符串返回索引和值。
	// 2, map返回键和值。
	// 3, 通道（channel）只返回通道内的值。
	d := "hello 循环"
	for i, v := range d {
		fmt.Printf("%d %c \n", i, v)
	}

	fmt.Println("=======switch case=======")
	// 使用switch语句可方便地对大量的值进行条件判断。
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}

	// Go语言规定每个switch只能有一个default分支。
	// 一个分支可以有多个值，多个case值中间使用英文逗号分隔。
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}

	// 使用表达式，这时候switch语句后面不需要再跟判断变量
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}

	// fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}

	fmt.Println("=========goto(跳转到指定标签)==========")
	// goto
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")

	multiplicationTable()
}

func multiplicationTable() {
	fmt.Println("===== 9*9乘法表 ======")
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", j, i, i*j)
		}
		fmt.Println()
	}
}
