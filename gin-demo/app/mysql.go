/*
 * @Author: zhengzhuang
 * @Date: 2021-07-17 10:20:12
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-17 11:47:38
 * @Description: mysql操作
 * @FilePath: /01-study/gin-demo/app/mysql.go
 */
package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// user_id, username, sex, email

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

// country, city, telcode

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	Telcode int    `db:"telcode"`
}

var db *sqlx.DB

// 初始化连接数据库
func Conn() {
	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	println("开始连接数据库=======")
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}
	fmt.Println("数据库连接成功")
	db = database
	// defer db.Close()
}

// 数据添加(增)
func AddData(c *gin.Context) {
	fmt.Println("--  开始向数据库添加数据  --")
	r, err := db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "张三", "man", "zheng960108@163.com")
	if err != nil {
		fmt.Println("exec failed1, ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed2, ", err)
		return
	}
	fmt.Println("添加 success:", id)
}

// 数据删除(删)
func DeleteData(c *gin.Context) {
	fmt.Println("--  开始向数据库删除数据  --")
	r, err := db.Exec("delete from person where user_id=?", 2)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	row, err := r.RowsAffected()
	if err != nil {
		fmt.Println("row failed, ", err)
	}
	fmt.Println("删除 success: ", row)
}

// 数据修改(改)
func UpdateData(c *gin.Context) {
	fmt.Println("-- 开始更新数据库的数据 --")
	r, err := db.Exec("update person set username=? where user_id=?", "哈哈", 3)
	if err != nil {
		fmt.Println("exec failed: ", err)
		return
	}
	row, err := r.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}
	fmt.Println("更新 success: ", row)
}

// 数据查找(找)
func SelectData(c *gin.Context) {
	fmt.Println("-- 开始查找数据库的数据 --")
	var person []Person
	err := db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 3)
	if err != nil {
		fmt.Println("exec failed: ", err)
	}
	fmt.Println("查找 success: ", person)
}
