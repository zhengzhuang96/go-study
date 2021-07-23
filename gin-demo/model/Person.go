// Author: zhengzhuang
// Date: 2021-07-20 14:58:52
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-21 21:46:05
// Description: In User Settings Edit
// FilePath: /01-study/gin-demo/model/person.go
package model

type Person struct {
	Id       int64  `db:"id"`
	UserId   int64  `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}
