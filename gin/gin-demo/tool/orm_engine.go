// Author: zhengzhuang
// Date: 2021-07-20 14:39:30
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-21 21:46:53
// Description: 数据库连接
// FilePath: /01-study/gin-demo/tool/orm_engine.go
package tool

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var DbEngine *Orm

type Orm struct {
	*sqlx.DB
}

// var schema = `
// CREATE TABLE person (
//     first_name text,
//     last_name text,
//     email text
// );

// CREATE TABLE place (
//     country text,
//     city text NULL,
//     telcode integer
// )`

// 数据库连接
func OrmEngine(cfg *Config) error {
	database := cfg.Database
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName
	engine, err := sqlx.Open("mysql", conn)
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return err
	}

	// engine.MustExec(schema)

	if err != nil {
		return err
	}

	orm := new(Orm)
	orm.DB = engine
	DbEngine = orm
	return nil
}
