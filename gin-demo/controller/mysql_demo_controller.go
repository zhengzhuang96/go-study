// Author: zhengzhuang
// Date: 2021-07-17 10:20:12
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-21 21:44:21
// Description: mysql操作
// FilePath: /01-study/gin-demo/controller/mysql_demo_controller.go
package controller

import (
	"gin-demo/service"
	"gin-demo/tool"

	"github.com/gin-gonic/gin"
)

type MysqlDemoController struct{}

func (mdc *MysqlDemoController) Router(engine *gin.Engine) {
	engine.GET("/addData", mdc.addData)
	engine.GET("/deleteData", mdc.deleteData)
	engine.GET("/updateData", mdc.UpdateData)
	engine.GET("/selectData", mdc.SelectData)
}

// 数据添加(增)
func (mdc *MysqlDemoController) addData(c *gin.Context) {
	ms := service.MysqlDemoService{}
	isAdd := ms.AddData()
	if isAdd {
		tool.Success(c, "添加成功")
		return
	}
	tool.Failed(c, "添加失败")
}

// 数据删除(删)
func (mdc *MysqlDemoController) deleteData(c *gin.Context) {
	ms := service.MysqlDemoService{}
	isDelete := ms.DeleteData()
	if isDelete {
		tool.Success(c, "删除成功")
		return
	}
	tool.Failed(c, "删除失败")
}

// 数据修改(改)
func (mdc *MysqlDemoController) UpdateData(c *gin.Context) {
	ms := service.MysqlDemoService{}
	isUpdate := ms.UpdateData()
	if isUpdate {
		tool.Success(c, "修改成功")
		return
	}
	tool.Failed(c, "修改失败")
}

// 数据查找(找)
func (mdc *MysqlDemoController) SelectData(c *gin.Context) {
	ms := service.MysqlDemoService{}
	selectInfo := ms.SelectData()
	if selectInfo != nil {
		tool.Success(c, map[string]interface{}{
			"userId":   selectInfo.UserId,
			"username": selectInfo.Username,
			"sex":      selectInfo.Sex,
			"email":    selectInfo.Email,
		})
		return
	}
	tool.Failed(c, "查找失败")
}
