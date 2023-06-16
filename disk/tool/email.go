package tool

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func Mail(context string, emailarr string) {
	e := email.NewEmail()
	e.From = "Jordan Wright <m17607466074@163.com>"
	e.To = []string{emailarr}
	e.Subject = "验证码"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte(fmt.Sprintf("<h1>你的验证码为%s!</h1>", context))
	err := e.SendWithTLS("smtp.163.com:587", smtp.PlainAuth("", "m17607466074@163.com", "HKZYGLMUUOCYKXIS", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	fmt.Println(err)
}
