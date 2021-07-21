// Author: zhengzhuang
// Date: 2021-07-20 09:41:54
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-21 21:46:44
// Description: In User Settings Edit
// FilePath: /01-study/gin-demo/tool/config.go

package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName     string         `json:"app_name"`
	AppMode     string         `json:"app_mode"`
	AppHost     string         `json:"app_host"`
	AppPort     string         `json:"app_port"`
	Database    DatabaseConfig `json:"database"`
	RedisConfig RedisConfig    `json:"redis_config"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`
	ShowSql  bool   `json:"showsql"`
}

type RedisConfig struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

var _cfg *Config = nil

func GetConfig() *Config {
	return _cfg
}

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decode := json.NewDecoder(reader)
	if err = decode.Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, nil
}
