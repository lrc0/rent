package message

import (
	"crypto/tls"
	"github.com/go-gomail/gomail"
	"gopkg.in/logger.v1"
)

//SendMail .
func SendMail(receiver, message, theme string) error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "313352050@qq.com", "") // 发件人
	m.SetHeader("To",                                  // 收件人
		m.FormatAddress(receiver, ""),
	)
	m.SetHeader("Subject", theme)   // 主题
	m.SetBody("text/html", message) // 正文

	d := gomail.NewPlainDialer("smtp.qq.com", 465, "313352050@qq.com", "mzqnjpznckkdbjjd") // 发送邮件服务器、端口、发件人账号、发件人密码
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
