// Author: zhengzhuang
// Date: 2021-07-28 09:26:23
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-28 09:26:24
// Description: In User Settings Edit
// FilePath: /go-study/算法/冒泡排序/main.go
package main

import "fmt"

// 由小到大排序
func SmallLarge(arr *[6]int) {
	temp := 0
	for j := 0; j < len(arr)-1; j++ {
		for i := 0; i < len(arr)-1-j; i++ { // 排序完一次，就排序剩下的，len-1次
			if arr[i] > arr[i+1] {
				temp = arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
			}
		}
	}
}

func main() {
	arrData := [6]int{1, 3, 9, 5, 4, 7}
	fmt.Println("原值:", arrData)

	SmallLarge(&arrData)
	fmt.Println("由小到大排序:", arrData)
}
