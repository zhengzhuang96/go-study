/*
 * @Author: zhengzhuang
 * @Date: 2021-07-20 14:58:52
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-21 10:45:02
 * @Description: In User Settings Edit
 * @FilePath: /01-study/gin-demo/model/person.go
 */
package model

type Person struct {
	Id       int64  `xorm:"pk autoincr" json:"id"`
	UserId   int64  `xorm:"int" json:"user_id"`
	Username string `xorm:"varchar(256)" json:"username"`
	Sex      string `xorm:"varchar(256)" json:"sex"`
	Email    string `xorm:"varchar(256)" json:"email"`
}
