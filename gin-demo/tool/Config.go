/*
 * @Author: zhengzhuang
 * @Date: 2021-07-20 09:41:54
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-20 14:54:44
 * @Description: In User Settings Edit
 * @FilePath: /01-study/gin-demo/tool/Config.go
 */
package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName  string         `json:"app_name"`
	AppMode  string         `json:"app_mode"`
	AppHost  string         `json:"app_host"`
	AppPort  string         `json:"app_port"`
	Database DatabaseConfig `json:database"`
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

var _cfg *Config = nil

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