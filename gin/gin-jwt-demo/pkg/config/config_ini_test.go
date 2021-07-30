// Author: zhengzhuang
// Date: 2021-07-30 14:35:28
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-30 14:35:29
// Description: In User Settings Edit
// FilePath: /go-study/gin/gin-jwt-demo/config/server_test.go
package config

import (
	"fmt"
	"testing"
)

func TestServer(t *testing.T) {
	ConfigIni = (&configIni{}).Load("../../config/app.ini").Init()
	fmt.Println(ConfigIni)
	if ConfigIni == nil {
		t.Fail()
	}
}
