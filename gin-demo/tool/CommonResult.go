/*
 * @Author: zhengzhuang
 * @Date: 2021-07-20 17:01:37
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-20 17:09:20
 * @Description: In User Settings Edit
 * @FilePath: /01-study/gin-demo/tool/CommonResult.go
 */
package tool

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS int = 0000 // 操作成功
	FAILED  int = 9999 // 操作失败
)

// 普通成功返回
func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"msg":  "请求成功",
		"data": v,
	})
}

// 操作失败返回内容
func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": FAILED,
		"msg":  v,
	})
}
