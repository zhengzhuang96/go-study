/*
 * @Author: zhengzhuang
 * @Date: 2021-07-15 17:14:45
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-15 17:32:24
 * @Description: defer 延迟调用
 * @FilePath: /函数/defer(延迟调用)/main.go
 */
package main

import "fmt"

// defer 特性
// 1.关键字 defer 用于注册延迟调用
// 2.这些调用直到 return 前才被执行。因此，可以用来做资源清理
// 3.多个 defer 语句，先进入后出的方式执行
// 4.defer 语句中的变量，在 defer 声明时就决定了

// defer 用途
// 1.关闭文件句柄
// 2.锁资源释放
// 3.数据库链接释放

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " closed")
}

// func Close(t Test) {
// 	t.Close()
// }

func main() {
	var whatever [5]struct{}

	// defer 按照先进后出的方式执行
	for i := range whatever {
		defer fmt.Println(i)
	}

	var evers [5]struct{}
	for i := range evers {
		defer func() { fmt.Println(i) }()
	}

	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, v := range ts {
		t2 := v
		defer t2.Close()
	}
}
