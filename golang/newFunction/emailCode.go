package newFunction

import (
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendEmain(mailTo []string, subject string, body string) error  {
	mailConn := map[string]string{
		"user": "这里需要自己加配置信息",
		"pass": "这里需要自己加配置信息",
		"host": "这里需要自己加配置信息",
		"port": "这里需要自己加配置信息",
	}
	port, _ := strconv.Atoi(mailConn["port"])	//转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], "lxs"))

	m.SetHeader("To", mailTo...)		//发送给多个用户
	m.SetHeader("Subject", subject)	//设置邮件主题
	m.SetBody("text/html", body)	//设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)


	return err

}