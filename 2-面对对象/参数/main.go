/*
 * @Author: zhengzhuang
 * @Date: 2021-07-15 16:16:15
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-15 16:36:19
 * @Description: In User Settings Edit
 * @FilePath: /函数/参数/main.go
 */
package main

import "fmt"

// 定义互相交换值的函数
func swap(x, y *int) {
	var temp int

	// 修改指针地址的值 指针会变化
	temp = *x // 保存x的值
	*x = *y   // 将y值赋值给x
	*y = temp // 将temp值赋值给y
}

func test(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}

func main() {
	var a, b int = 1, 2
	/*
		调用 swap() 函数
		&a 指向 a 指针，a 变量的地址
		&b 指向 b 指针，b 变量的地址
	*/
	swap(&a, &b)

	// fmt.Println(a, b)

	s := []int{1, 2, 3}

	println(test("sum: %d", s...))

}
