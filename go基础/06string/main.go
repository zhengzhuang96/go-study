/*
 * @Author: zhengzhuang
 * @Date: 2021-07-14 11:23:31
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-14 11:54:17
 * @Description: 字符串
 * @FilePath: /day01/06string/main.go
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "hello"
	s2 := "你好"
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println("-------------------------")

	// 字符串转义符
	myStr := "abc\r\n123\r\n"
	fmt.Println(myStr)

	fmt.Println("-------------------------")

	// 多行字符串,	反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。
	s3 := `第一行
	第二行
	第三行`
	fmt.Println(s3)

	fmt.Println("-------------------------")

	// 字符串的常用操作
	fmt.Println(len("hello"))                                  // 求长度
	fmt.Println("字符串拼接：", "a"+"b")                             // 拼接字符串
	fmt.Println("字符串拼接：", fmt.Sprintf("e=%s&a=%s", "c", "d"))  // 拼接字符串
	fmt.Println("切割字符串：", strings.Split("strings", "i"))       // 切割字符串
	fmt.Println("判断是否包含：", strings.Contains("strings", "ta"))  // 判断是否包含	false
	fmt.Println("判断前缀：", strings.HasPrefix("hello", "he"))     // 判断前缀	true
	fmt.Println("判断后缀：", strings.HasSuffix("hello", "lo"))     // 判断后缀	true
	fmt.Println("子串首次出现的位置：", strings.Index("abbbbc", "b"))    // 子串首次出现的位置
	fmt.Println("子串最后出现的位置：", strings.LastIndex("abbbc", "b")) // 子串最后出现的位置
	var s = []string{"11", "22", "33"}
	fmt.Println("字符串拼接：", strings.Join(s, "|")) // 将s中的子串连接成一个单独的字符串，子串之间用｜分隔  11|22|33
}
