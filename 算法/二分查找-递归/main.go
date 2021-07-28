// Author: zhengzhuang
// Date: 2021-07-28 10:55:19
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-28 10:55:20
// Description: In User Settings Edit
// FilePath: /go-study/算法/二分查找-递归/main.GO
package main

import "fmt"

//  使用二分查找，数组或切片必须是有序的

// 二分查找+递归
func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {

	// 判断leftIndex 是否大于 rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}

	// 先找到中间的下标
	middle := (leftIndex + rightIndex) / 2

	if arr[middle] > findVal {
		// 说明我们要查找的数，应该在 leftIndex --- middel-1 之间
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if arr[middle] < findVal {
		// 说明我们要查找的数，应该在 middel+1 --- rightIndex
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		// 找到了
		fmt.Println("找到了，下标为:", middle)
	}
}

func main() {
	arr := [6]int{1, 8, 10, 89, 1000, 1345}
	BinaryFind(&arr, 0, len(arr)-1, 89)
}
