// Author: zhengzhuang
// Date: 2021-07-20 15:30:41
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-21 17:57:43
// Description: In User Settings Edit
// FilePath: /01-study/gin-demo/service/mysql_demo_service.go
package service

import (
	"gin-demo/db"
	"gin-demo/model"
	"gin-demo/tool"
)

type MysqlDemoService struct{}

func (mds *MysqlDemoService) AddData() bool {
	r := model.Person{UserId: 2, Username: "我的名字", Sex: "男", Email: "121730414@qq.com"}
	personDb := db.PersonDb{Orm: tool.DbEngine}
	result := personDb.InsertData(r)
	return result > 0
}

func (mds *MysqlDemoService) DeleteData() bool {
	r := model.Person{UserId: 2}
	personDb := db.PersonDb{Orm: tool.DbEngine}
	result := personDb.DeleteData(r)
	return result > 0
}

func (md *MysqlDemoService) UpdateData() bool {
	r := model.Person{UserId: 1, Username: "哈哈哈"}
	personDb := db.PersonDb{Orm: tool.DbEngine}
	result := personDb.UpdateData(r)
	return result
}

func (md *MysqlDemoService) QueryRowData() *model.Person {
	r := model.Person{UserId: 1}
	personDb := db.PersonDb{Orm: tool.DbEngine}
	return personDb.QueryRowData(r)
}
