/*
 * @Author: zhengzhuang
 * @Date: 2021-07-20 15:30:41
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-20 16:52:49
 * @Description: In User Settings Edit
 * @FilePath: /01-study/gin-demo/service/MysqlDemoService.go
 */
package service

import (
	"gin-demo/dao"
	"gin-demo/model"
	"gin-demo/tool"
)

type MysqlDemoService struct{}

func (mds *MysqlDemoService) AddData() bool {
	r := model.Person{UserId: 2, Username: "username", Sex: "男", Email: "121730414@qq.com"}
	personDao := dao.PersonDao{Orm: tool.DbEngine}
	result := personDao.InsertData(r)
	return result > 0
}

func (mds *MysqlDemoService) DeleteData() bool {
	r := model.Person{UserId: 2}
	personDao := dao.PersonDao{Orm: tool.DbEngine}
	result := personDao.DeleteData(r)
	return result > 0
}

func (md *MysqlDemoService) UpdateData() bool {
	r := model.Person{UserId: 1, Username: "哈哈哈"}
	personDao := dao.PersonDao{Orm: tool.DbEngine}
	result := personDao.UpdateData(r)
	return result > 0
}

func (md *MysqlDemoService) SelectData() *model.Person {
	r := model.Person{UserId: 1}
	personDao := dao.PersonDao{Orm: tool.DbEngine}
	// result := personDao.SelectData(r)
	return personDao.SelectData(r)
}
