// Author: zhengzhuang
// Date: 2021-07-24 14:42:12
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-24 14:42:12
// Description: In User Settings Edit
// FilePath: /01-study/im-system-demo/server.go

package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port string

	// 在线用户的列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan string
}

// 创建一个server的接口
func NewServer(ip string, port string) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 启动服务器的接口
func (s *Server) Start() {
	// socket listen
	lister, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// close listen socket
	defer lister.Close()

	// 启动监听Message的goroutine
	go s.ListenMessager()

	for {
		// accept
		conn, err := lister.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}
		// do handler
		go s.Handler(conn)
	}
}

// 监听Message广播消息channel的goroutine，一旦有消息就发送给全部的在线User
func (s *Server) ListenMessager() {
	for {
		msg := <-s.Message

		// 将msg发送给全部的在线User
		s.mapLock.Lock()
		for _, v := range s.OnlineMap {
			v.C <- msg
		}
		s.mapLock.Unlock()
	}
}

func (s *Server) Handler(conn net.Conn) {
	// 当前链接的业务
	// fmt.Println("链接建立成功")

	user := NewUser(conn, s)

	// 用户上线, 将用户加入到onlineMap中
	user.online()

	// 监听用户是否活跃的channel
	isLive := make(chan bool)

	// 接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}
			if n == 0 {
				user.Offline()
				return
			}

			// 提取用户的消息(去除'\n')
			msg := string(buf[:n-1])

			// 用户针对msg进行处理
			user.DoMessage(msg)

			// 用户的任意消息，代表当前用户是一个活跃的
			isLive <- true
		}
	}()

	// 当前handler阻塞
	for {
		select {
		case <-isLive:
			// 当前用户是活跃的，应该重置定时器
			// 不做任何事情，为了激活select，更新下面的定时器
		case <-time.After(time.Second * 300):
			// 已经超时
			// 将当前的User强制的关闭

			user.SendMsg("你被提了")

			// 销毁用的资源
			close(user.C)

			// 关闭链接
			conn.Close()

			// 退出当前handler
			// runtime.Goexit() // return
			return
		}
	}
}

// 广播消息的方法
func (s *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	s.Message <- sendMsg
}
