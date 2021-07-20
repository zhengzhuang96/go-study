/*
 * @Author: zhengzhuang
 * @Date: 2021-07-19 17:42:44
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-19 17:46:46
 * @Description: In User Settings Edit
 * @FilePath: /01-study/gin-demo/service/testService.go
 */

package Services

import (
	"gin/Models"
)

type Test struct {
	Id      int    `json:"id"`
	Testcol string `json:"testcol"`
}

func (this *Test) Insert() (id int, err error) {
	var testModel Models.Test
	testModel.Id = this.Id
	testModel.Testcol = this.Testcol
	id, err = testModel.Insert()
	return
}
