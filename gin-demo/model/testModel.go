/*
 * @Author: zhengzhuang
 * @Date: 2021-07-19 17:43:07
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-19 17:46:37
 * @Description: In User Settings Edit
 * @FilePath: /01-study/gin-demo/model/testModel.go
 */
package model

type Test struct {
	Id      int
	Testcol string `gorm:"column:testcol"`
}

// 设置Test的表名为`test`
// func (Test) TableName() string {
//     return "test"
// }

func (this *Test) Insert() (id int, err error) {
	result := mysql.DB.Create(&this)
	id = this.Id
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
