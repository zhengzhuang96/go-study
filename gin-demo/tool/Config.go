/*
 * @Author: zhengzhuang
 * @Date: 2021-07-20 09:41:54
 * @LastEditors: zhengzhuang
 * @LastEditTime: 2021-07-20 09:51:42
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
	AppName string `json:"app_name"`
	AppMode string `json:"app_mode"`
	AppHost string `json:"app_host"`
	AppPort string `json:"app_port"`
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
