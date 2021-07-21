// Author: zhengzhuang
// Date: 2021-07-16 11:07:28
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-21 21:44:39
// Description: API参数
// FilePath: /01-study/gin-demo/controller/parameters_controller.go
package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ParametersController struct{}

func (parameters *ParametersController) Router(engine *gin.Engine) {
	engine.GET("/apiParameters/:name/*action", parameters.apiParameters)
	engine.GET("/urlParameters", parameters.urlParameters)
}

// 可以通过Context的Param方法来获取API参数
// localhost:8000/xxx/zhangsan	来获取路由后的值
func (parameters *ParametersController) apiParameters(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	fmt.Printf("name：%s, action: %s", name, action)
	action = strings.Trim(action, "/")
	c.String(http.StatusOK, fmt.Sprintf("name %s + action %s", name, action))
}

// URL参数	来获取?后的值
// URL参数可以通过DefaultQuery()或Query()方法获取
// DefaultQuery()若参数不村则，返回默认值，Query()若不存在，返回空串
// API?name=zs
func (parameters *ParametersController) urlParameters(c *gin.Context) {
	name := c.DefaultQuery("name", "默认值")
	path := c.DefaultQuery("path", "默认值")
	c.String(http.StatusOK, fmt.Sprintf("Hello %s, path %s", name, path))
}
