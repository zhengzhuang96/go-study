// Author: zhengzhuang
// Date: 2021-07-26 15:15:45
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-26 15:15:45
// Description: fmt
// FilePath: /go-study/package/fmt/main.go

package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	// 通用
	person := Person{Name: "亚索"}
	fmt.Printf("%v\n", person)  // 默认格式输出	{亚索}
	fmt.Printf("%+v\n", person) // 添加字段名输出	{Name:亚索}
	fmt.Printf("%#v\n", person) // 值的go语法表示 main.Person{Name:"亚索"}
	fmt.Printf("%T\n", person)  // 类型	main.Person

	// 布尔值
	fmt.Printf("%t\n", true)  // 单词true
	fmt.Printf("%t\n", false) // 单词false

	// 整数
	fmt.Printf("%b\n", 5)      // 二进制	101
	fmt.Printf("%c\n", 0x4E2d) // 对应的unicode码值 中
	fmt.Printf("%d\n", 0x12)   // 十进制 18
	fmt.Printf("%o\n", 10)     // 八进制 12
	fmt.Printf("%q\n", 0x4E2d) // 该值对应的单引号括起来的语法字面量	'中'
	fmt.Printf("%x\n", 13)     // 十六进制，a-f d
	fmt.Printf("%X\n", 13)     // 十六进制，A-F D
	fmt.Printf("%U\n", 0x4E2d) // Unicode格式	U+4E2d

	// 浮点数与复数
	fmt.Printf("%b\n", 10.20) // 无小数部分、二进制指数的科学计数法	5742089524897382p-49
	fmt.Printf("%e\n", 10.20) // 科学计数法	1.020000e+01
	fmt.Printf("%E\n", 10.20) // 科学计数法 1.020000E+01
	fmt.Printf("%f\n", 10.20) // 有小数部分但无指数部分 10.200000
	fmt.Printf("%F\n", 10.20) // 等价于%f 10.200000
	fmt.Printf("%g\n", 10.20) // 根据实际情况采用%e或%f格式 10.2
	fmt.Printf("%G\n", 10.20) // 根据实际情况采用%E或%F格式 10.2

	// 字符串和[]byte
	fmt.Printf("%s\n", []byte("Go语言")) // 直接输出字符串或者[]byte Go语言
	fmt.Printf("%q\n", "Go语言")         // 该值对应的双引号括起来的Go语法字符串字面值 "Go语言"
	fmt.Printf("%x\n", "golang")       // 每个字节用两字符十六进制数表示（使用a-f） 676f6c616e67
	fmt.Printf("%X\n", "golang")       // 每个字节用两字符十六进制数表示（使用A-F） 676F6C616E67

	// 指针
	fmt.Printf("%p\n", &person) // 表示为十六进制，并加上前导的0x	0xc000096220
}
