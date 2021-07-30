// Author: zhengzhuang
// Date: 2021-07-30 14:06:44
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-30 14:06:47
// Description: In User Settings Edit
// FilePath: /go-study/gin/gin-jwt-demo/tool/Ini.go
package config

import (
	"github.com/go-ini/ini"
)

var ConfigIni *configIni

type configIni struct {
	Secret string
	source *ini.File
}

func (s *configIni) Load(path string) *configIni {
	var err error

	s.source, err = ini.Load(path)

	if err != nil {
		panic(err)
	}
	return s
}

func (s *configIni) Init() *configIni {
	//判断配置是否加载成功
	if s.source == nil {
		return s
	}
	s.Secret = s.source.Section("tokenSecret").Key("secret").MustString("123admin123")
	return s
}
