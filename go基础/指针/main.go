/*
 * @Author: zhengzhuang
 * @Date: 2021-07-14 17:10:01
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-15 12:00:21
 * @Description: 指针
 * @FilePath: /go基础/指针/main.go
 */
package main

import "fmt"

// 指针
// &（取地址）和*（根据地址取值）
func main() {
	// a := 10
	// b := &a // 取变量a的地址，将指针保存到b中
	// fmt.Printf("a:%d ptr:%p\n", a, &a)
	// fmt.Printf("a:%p type:%T\n", b, b)
	// fmt.Println(&b) // 0xc00000e018
	// c := *b         // 指针取值（根据指针去内存取值）
	// fmt.Printf("type of c:%T\n", c)
	// fmt.Printf("value of c:%v\n", c)

	// a := 10
	// modify1(a)
	// fmt.Println(a) // 10
	// modify2(&a)
	// fmt.Println(a) // 100

	// var p *string
	// fmt.Println(p)
	// fmt.Printf("p的值是%v\n", p)

	// if p != nil {
	// 	fmt.Println("非空")
	// } else {
	// 	fmt.Println("空值")
	// }

	// a := new(int)
	// b := new(bool)
	// fmt.Printf("%T\n", a) // *int
	// fmt.Printf("%T\n", b) // *bool
	// fmt.Println(*a)       // 0
	// fmt.Println(*b)       // false

	var b map[string]int
	b = make(map[string]int, 10)
	b["测试"] = 100
	fmt.Println(b)

	var a int
	fmt.Println(&a)
	var p *int
	p = &a
	*p = 20
	fmt.Println(a)
}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}
