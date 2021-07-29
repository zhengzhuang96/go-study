// Author: zhengzhuang
// Date: 2021-07-28 15:58:45
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-28 15:59:57
// Description: In User Settings Edit
// FilePath: /go-study/go-web-demo/cgo/RouterHandler.go
package cgo

import (
	"fmt"
	"net/http"
	"strings"
)

var Router *RouterHandler = new(RouterHandler)

type RouterHandler struct {
}

var mux = make(map[string]func(http.ResponseWriter, *http.Request))

func (rh *RouterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if fun, ok := mux[r.URL.Path]; ok {
		fun(w, r)
		return
	}
	//静态资源
	if strings.HasPrefix(r.URL.Path, "/img") {
		if fun, ok := mux["/img"]; ok {
			fun(w, r)
			return
		}
	}
	http.Error(w, "error URL:"+r.URL.String(), http.StatusBadRequest)
}

func (p *RouterHandler) Router(relativePath string, handler func(http.ResponseWriter, *http.Request)) {
	mux[relativePath] = handler
}

func (p *RouterHandler) POST(relativePath string, handler func(http.ResponseWriter, *http.Request)) {
	mux[relativePath] = handler
	mux[http.MethodGet] = handler
}
