/*
 * @Author: zhengzhuang
 * @Date: 2021-07-14 15:24:40
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-14 15:55:32
 * @Description: In User Settings Edit
 * @FilePath: /day01/10数组/main.go
 */
package main

import "fmt"

// 数组定义
// var 数组变量名 [元素数量]T

func main() {
	// 数组
	// 定义一个长度为3元素类型为int的数组a
	var a [3]int
	fmt.Println(a)

	// var a [5]int， 数组的长度必须是常量，并且长度是数组类型的一部分。一旦定义，长度不能变。 [5]int和[10]int是不同的类型。
	// var b [2]int
	// var c [3]int
	// b = c //不可以这样做，因为此时a和b是不同的类型

	// 数组的初始化
	// var testArray [3]int                        // 数组会初始化为int类型的零值
	// var numArray = [3]int{1, 2}                 // 使用指定的初始值完成初始化
	// var cityArray = [3]string{"北京", "上海", "深圳"} // 使用指定的初始值完成初始化
	// fmt.Println(testArray)                      // [0 0 0]
	// fmt.Println(numArray)                       // [1 2 0]
	// fmt.Println(cityArray)                      // [北京 上海 深圳]

	// 可以让编译器根据初始值的个数自行推断数组的长度
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)                          //[0 0 0]
	fmt.Println(numArray)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
	fmt.Println(cityArray)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string

	// 使用指定索引值的方式来初始化数组
	c := [...]int{1: 1, 3: 5}
	fmt.Println(c)                  // [0 1 0 5]		因为制定了第1个和第3个，所以第0个和第3个是默认0
	fmt.Printf("type of c:%T\n", c) // type of c:[4]int

	// 数组的遍历
	fmt.Println("====================")
	var d = [...]string{"北京", "上海", "深圳"}
	// 方法1: for循环便利
	for i := 0; i < len(d); i++ {
		fmt.Println(d[i])
	}

	// 方法2：for range遍历
	for index, v := range d {
		fmt.Println(index, v)
	}

	// 多维数组
	// 二维数组的定义
	fmt.Println("========二维数组========")
	f := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(f)
	fmt.Println(f[2][1])

	for _, v := range f {
		for _, v2 := range v {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}

	// 多维数组只有第一层可以使用...来让编译器推导数组长度
	// 支持的写法
	g := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	// 不支持多维数组的内层使用...
	// h := [3][...]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }
	fmt.Println(g)

	// 数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
	h := [3]int{10, 20, 30}
	modifyArray(h) //在modify中修改的是a的副本x
	fmt.Println(h) //[10 20 30]
	l := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(l) //在modify中修改的是b的副本x
	fmt.Println(l)  //[[1 1] [1 1] [1 1]]
}

func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}
