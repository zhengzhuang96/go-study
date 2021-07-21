// Author: zhengzhuang
// Date: 2021-07-17 11:56:16
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-21 21:44:48
// Description: redis操作
// FilePath: /01-study/gin-demo/controller/redis_controller.go
package controller

import (
	"gin-demo/tool"
	"time"

	"github.com/gin-gonic/gin"
)

type RedisController struct{}

func (rc *RedisController) Router(engine *gin.Engine) {
	engine.GET("/redis/setRedis", rc.setRedis)
	engine.GET("/redis/getRedis", rc.getRedis)
	engine.GET("/redis/batchRedis", rc.batchRedis)
}

// 设置一个值, 并设置一个时间10s
func (rc *RedisController) setRedis(c *gin.Context) {
	tool.RedisStore.Set("name", "我是存储的值", 10*time.Second)
	tool.Success(c, "存储成功")
}

// 获取一个值或多个值
func (rc *RedisController) getRedis(c *gin.Context) {
	dataInfo := tool.RedisStore.Get("name")
	if len(dataInfo) == 0 {
		tool.Failed(c, "没有找到")
		return
	}
	tool.Success(c, dataInfo)
}

// 批量存储到redis
func (rc *RedisController) batchRedis(c *gin.Context) {
	err := tool.RedisStore.MSet("name1", "我是name1的值", "name2", "我是name2的值")
	if err != nil {
		tool.Failed(c, "存储redis失败")
		return
	}
	tool.Success(c, "存储redis成功")
}
