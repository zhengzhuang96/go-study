/*
 * @Author: zhengzhuang
 * @Date: 2021-07-14 14:11:35
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-14 14:17:49
 * @Description: if
 * @FilePath: /day01/08if/main.go
 */
package main

import "fmt"

// Go语言中if条件判断的格式如下：
// if 表达式1 {
// 	分支1
// } else if 表达式2 {
// 	分支2
// } else{
// 	分支3
// }

func main() {

	// 正常写法
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// 特殊写法，根据变量值进行判断	作用域，num变量此时只在判断语句中生效
	if num := 65; num >= 90 {
		fmt.Println("A")
	} else if num > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

}
