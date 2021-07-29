// Author: zhengzhuang
// Date: 2021-07-29 08:46:12
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-29 08:49:03
// Description: In User Settings Edit
// FilePath: /go-study/go-web/cgo/engine.go

package cgo

import (
	"net/http"
)

//HandlerFunc定义封装http使用的请求处理程序
type HandlerFunc func(*Context)

// 引擎是所有请求的uni处理程序
type Engine struct {
	router *router
}

// New是封装http Engine的构造函数
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET定义添加GET请求的方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST定义添加POST请求的方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run定义启动http服务器的方法
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
