/*
 * @Author: zhengzhuang
 * @Date: 2021-07-20 15:12:18
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-20 16:51:37
 * @Description: In User Settings Edit
 * @FilePath: /01-study/gin-demo/dao/PersonDao.go
 */
package dao

import (
	"gin-demo/model"
	"gin-demo/tool"
)

type PersonDao struct {
	*tool.Orm
}

func (pd *PersonDao) InsertData(person model.Person) int64 {
	result, err := pd.InsertOne(&person)
	if err != nil {
		println(err.Error())
	}
	return result
}

func (pd *PersonDao) DeleteData(person model.Person) int64 {
	result, err := pd.Where("user_id=?", person.UserId).Delete(&person)
	if err != nil {
		println(err.Error())
	}
	return result
}

func (pd *PersonDao) UpdateData(person model.Person) int64 {
	result, err := pd.Where("user_id=?", person.UserId).Update(&person)
	if err != nil {
		println(err.Error())
	}
	return result
}

// func (pd *PersonDao) SelectData(person model.Person) map[string]string {
func (pd *PersonDao) SelectData(person model.Person) *model.Person {
	var personMap model.Person
	// result, err := pd.Where("user_id=?", person.UserId).Get(&person)
	// result, err := pd.Where("user_id=?", person.UserId).Find(&person)
	_, err := pd.Orm.Where("user_id=?", person.UserId).Get(&personMap)
	if err != nil {
		panic(err.Error())
	}
	return &personMap
	// return result
}
