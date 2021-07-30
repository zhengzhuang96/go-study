// Author: zhengzhuang
// Date: 2021-07-30 15:07:11
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-30 15:07:14
// Description: In User Settings Edit
// FilePath: /go-study/gin/gin-jwt-demo/pkg/config/init.go

package config

func init() {
	ConfigIni = (&configIni{}).Load("config/app.ini").Init()
}
