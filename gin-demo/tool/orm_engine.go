// Author: zhengzhuang
// Date: 2021-07-20 14:39:30
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-21 21:46:53
// Description: 数据库连接
// FilePath: /01-study/gin-demo/tool/orm_engine.go
package tool

import (
	"gin-demo/model"

	"github.com/go-xorm/xorm"
)

var DbEngine *Orm

type Orm struct {
	*xorm.Engine
}

// 数据库连接
func OrmEngine(cfg *Config) (*Orm, error) {
	database := cfg.Database
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.Charset
	engine, err := xorm.NewEngine("mysql", conn)
	if err != nil {
		return nil, err
	}

	engine.ShowSQL(database.ShowSql)

	err = engine.Sync2(new(model.Person))
	if err != nil {
		return nil, err
	}

	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm
	return orm, nil
}
