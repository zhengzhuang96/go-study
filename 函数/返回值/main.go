/*
 * @Author: zhengzhuang
 * @Date: 2021-07-15 16:42:16
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-15 17:13:25
 * @Description: 返回值
 * @FilePath: /函数/返回值/main.go
 */

package main

import "fmt"

func add(a, b int) (c int) {
	// 命名返回参数允许 defer 延迟调用通过闭包读取和修改 返回值是等带这些调用直到 return 前才被执
	defer func() {
		c += 100
	}()

	c = a + b
	return
}

func calc(a, b int) (sum int, avg int) {
	sum = a + b
	avg = (a + b) / 2
	return
}

func test() (int, int) {
	return 1, 2
}

func adds(x, y int) int {
	return x + y
}

func sumFun(n ...int) int {
	var x int
	for _, i := range n {
		x += i
	}
	return x
}

func main() {
	var a, b int = 1, 2
	c := add(a, b)
	fmt.Println(c)
	fmt.Println("===========")
	// Golang返回值不能用容器对象接收多返回值。只能用多个变量，或 "_" 忽略。
	sum, avg := calc(a, b)
	fmt.Println(a, b, c, sum, avg)

	println(add(test()))
	println(sumFun(test()))
}
