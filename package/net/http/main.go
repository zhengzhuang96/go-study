// Author: zhengzhuang
// Date: 2021-07-26 11:19:59
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-26 11:20:06
// Description: go/http 包
// FilePath: /go-study/package/net/http/main.go

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	client := &NetHttp{}
	// client.HttpGet()
	client.HttpPost()
}

type NetHttp struct {
}

// get请求
func (nh *NetHttp) HttpGet() {
	resp, err := http.Get("https://wwww.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.StatusCode)
}

// post请求
func (nh *NetHttp) HttpPost() {
	requestUrl := "http://www.baidu.com/"
	postValue := url.Values{
		"username": {"hangmeimei"},
		"address":  {"anhui"},
		"subject":  {"world"},
		"form":     {"beij"},
	}
	//request, err := http.PostForm(requestUrl, postValue)

	body := bytes.NewBufferString(postValue.Encode())
	request, err := http.Post(requestUrl, "text/html", body)
	if err != nil {
		fmt.Println(err)
	}

	defer request.Body.Close()
	fmt.Println(request.StatusCode)
	if request.StatusCode == 200 {
		rb, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Println(rb)
		}
		fmt.Println(string(rb))
	}
}
