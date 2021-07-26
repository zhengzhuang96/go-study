// Author: zhengzhuang
// Date: 2021-07-26 14:04:48
// LastEditors: zhengzhuang
// LastEditTime: 2021-07-26 14:04:49
// Description: net/smtp 邮件
// FilePath: /go-study/package/net/smtp/main.go
package main

import (
	"fmt"
	"net/smtp"
)

type SmtpMail struct {
	SMTPHost     string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
}

func (sm *SmtpMail) SendEmail(receiver string) {
	//PlainAuth返回实现普通身份验证的身份验证
	//RFC4616中定义的机制。返回的身份验证使用给定的
	//对主机进行身份验证并充当身份的用户名和密码。
	//通常标识应该是空字符串，作为用户名。
	//只有当连接使用TLS时，PlainAuth才会发送凭据
	//或连接到本地主机。否则身份验证将失败
	auth := smtp.PlainAuth("", sm.SMTPUsername, sm.SMTPPassword, sm.SMTPHost)
	msg := []byte("Subject: 这里是标题内容\r\n\r\n" + "这里是正文内容\r\n")
	err := smtp.SendMail(sm.SMTPHost+sm.SMTPPort, auth, sm.SMTPUsername, []string{receiver}, msg)
	if err != nil {
		fmt.Println("failed to send email:", err)
	}
}

func main() {
	mail := &SmtpMail{
		SMTPHost:     "imap.qq.com",
		SMTPPort:     ":25",
		SMTPUsername: "121730414@qq.com",
		SMTPPassword: "gqdhpzsmpqpbbgbd",
	}
	mail.SendEmail("zheng960108@163.com")
}
