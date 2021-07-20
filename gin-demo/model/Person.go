/*
 * @Author: zhengzhuang
 * @Date: 2021-07-20 14:58:52
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-20 16:18:31
 * @Description: In User Settings Edit
 * @FilePath: /01-study/gin-demo/model/Person.go
 */
package model

type Person struct {
	UserId   int    `xorm:"varchar(11)" json:"user_id"`
	Username string `xorm:"varchar(256)" json:"username"`
	Sex      string `xorm:"varchar(256)" json:"sex"`
	Email    string `xorm:"varchar(256)" json:"email"`
}
