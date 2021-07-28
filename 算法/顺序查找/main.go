// Author: zhengzhuang
// Date: 2021-07-28 10:08:55
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-28 10:08:55
// Description: In User Settings Edit
// FilePath: /go-study/算法/顺序查找/main.go
package main

import "fmt"

func main() {
	arr := [6]string{"刘备", "关羽", "张飞", "曹操", "孙权", "诸葛亮"}
	var search string

	fmt.Println("请输入要查找的人物...")
	fmt.Scanln(&search)

	// 一种写法
	// for i := 0; i < len(arr); i++ {
	// 	if search == arr[i] {
	// 		fmt.Printf("查到了 %v，下标为 %v\n", search, i)
	// 	} else if i == len(arr)-1 {
	// 		fmt.Println("没有查到")
	// 	}
	// }

	// 另一种写法
	index := -1
	for i := 0; i < len(arr); i++ {
		if search == arr[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("查到了 %v, 下标为 %v\n", search, index)
	} else {
		fmt.Println("没有查到")
	}
}
