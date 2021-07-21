/*
 * @Author: zhengzhuang
 * @Date: 2021-07-21 09:40:27
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-21 14:47:29
 * @Description: In User Settings Edit
 * @FilePath: /01-study/gin-demo/tool/redis_store.go
 */
package tool

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type RedisStores struct {
	client *redis.Client
}

var RedisStore RedisStores

func InitRedisStore() *RedisStores {
	config := GetConfig().RedisConfig
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr + ":" + config.Port,
		Password: config.Password,
		DB:       config.Db,
	})
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("redis连接失败:", err.Error())
	}

	RedisStore = RedisStores{client: client}

	return &RedisStore
}

// set
func (rs *RedisStores) Set(id string, value string, expiration time.Duration) {
	err := rs.client.Set(id, value, expiration).Err()
	if err != nil {
		log.Println(err)
	}
}

// set
func (rs *RedisStores) MSet(pairs ...interface{}) error {
	err := rs.client.MSet(pairs...).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// get
func (rs *RedisStores) Get(keys ...string) []interface{} {
	val, err := rs.client.MGet(keys...).Result()
	if err != nil {
		log.Println(err)
	}
	return val
}

// del
func (rs *RedisStores) Clear(id string) string {
	err := rs.client.Del(id).Err()
	if err != nil {
		log.Println(err)
		return ""
	}
	return ""
}
