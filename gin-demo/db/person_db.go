// Author: zhengzhuang
// Date: 2021-07-20 15:12:18
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-21 21:45:14
// Description: In User Settings Edit
// FilePath: /01-study/gin-demo/dao/person_dao.go
package db

import (
	"fmt"
	"gin-demo/model"
	"gin-demo/tool"
)

type PersonDb struct {
	*tool.Orm
}

func (db *PersonDb) InsertData(person model.Person) int64 {
	sqlStr := "INSERT into person(user_id, username, sex, email) values (?,?,?,?)"
	ret, err := db.Exec(sqlStr, 123123, "小王子", "男", "zheng960108@163.com")
	if err != nil {
		println(err.Error())
		return 0
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return 0
	}
	return theID
}

func (db *PersonDb) DeleteData(person model.Person) int64 {
	sqlStr := "DELETE from person where id = ?"
	result, err := db.Exec(sqlStr, 6)
	if err != nil {
		println(err.Error())
	}
	n, err := result.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
	}
	return n
}

func (db *PersonDb) UpdateData(person model.Person) bool {
	sqlStr := "update person set username = ? where id = ?"
	_, err := db.Exec(sqlStr, "新名字", 7)
	if err != nil {
		println(err.Error())
		return false
	}
	return true
}

func (db *PersonDb) QueryRowData(person model.Person) *model.Person {
	var personMap model.Person
	sqlStr := "select id, user_id, username, sex, email from person where id = ?"
	err := db.Get(&personMap, sqlStr, 7)
	if err != nil {
		panic(err.Error())
	}
	return &personMap
}
