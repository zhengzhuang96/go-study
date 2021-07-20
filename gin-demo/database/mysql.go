/*
 * @Author: zhengzhuang
 * @Date: 2021-07-19 17:44:41
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-19 18:18:28
 * @Description: In User Settings Edit
 * @FilePath: /01-study/gin-demo/database/mysql.go
 */
package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

// 初始化连接数据库
func Conn() {
	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	println("开始连接数据库=======")
	data, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}
	fmt.Println("数据库连接成功")
	DB = data
	// defer db.Close()
}

// 创建数据
func Create() {

}
