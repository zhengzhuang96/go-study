// Author: zhengzhuang
// Date: 2021-07-28 15:54:25
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-28 15:54:25
// Description: In User Settings Edit
// FilePath: /go-study/go-web-demo/controller/test_controller.go
package controller

import (
	"fmt"
	"go-web-demo/cgo"
	"net/http"
)

type TextController struct{}

func (tc *TextController) Router(r *cgo.RouterHandler) {
	r.Router("/getText", tc.getText)
}

func (tc *TextController) getText(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("你好")
	// fmt.Fprintf(w, "abc")
	fmt.Fprintf(w, "http://localhost:8080/%s", "key")
}
