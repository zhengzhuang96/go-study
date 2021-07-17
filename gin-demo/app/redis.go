/*
 * @Author: zhengzhuang
 * @Date: 2021-07-17 11:56:16
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-17 16:03:36
 * @Description: redis操作
 * @FilePath: /01-study/gin-demo/app/redis.go
 */
package app

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

// 连接Redis，并设置一个abc的值是100
func ConnectRedis(c *gin.Context) {
	r, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	fmt.Println("redis conn success")
	// 延迟调用，defer特点：等函数一切程序跑完在调用
	defer r.Close()

	// 设置一个abc赋值成100
	_, err = r.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 从redis中取出存储的一个值
func GetRedis(c *gin.Context) {
	// 先链接
	r, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	fmt.Println("redis conn success")
	// 延迟调用，defer特点：等函数一切程序跑完在调用
	defer r.Close()

	// 获取abc的值
	g, err := redis.Int(r.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc failed", err)
		return
	}
	fmt.Println("abc的值是：", g)
}

// 批量存储到redis
func BatchRedis(c *gin.Context) {
	// 先连接redis
	r, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	fmt.Println("redis conn success")
	// 延时调用，defer特点：等函数里面一切程序跑完在调用
	defer r.Close()

	// 存储多个值
	_, err = r.Do("MSet", "a", 100, "b", "200")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("redis 存储完成")
}

// 取多个存储在redis中的值
func GetBatchRedis(c *gin.Context) {
	// 先连接redis
	r, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	fmt.Println("redis conn success")
	defer r.Close()

	// 获取多个存储的值
	t, err := redis.Ints(r.Do("MGet", "a", "b"))
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range t {
		fmt.Println(v)
	}
}

// 设置一个过期时间, 没成功
func ExpireTimeRedis(c *gin.Context) {
	// 连接redis
	r, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Close()

	_, err = r.Do("expire", "abcd", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("存储成功___")
}
