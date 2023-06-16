package main

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func main() {
	for {

		e := email.NewEmail()
		e.From = "Jordan Wright <m17607466074@163.com>"
		e.To = []string{"2110524485@qq.com"}
		e.Subject = "勺勺猪"
		e.Text = []byte("Text Body is, of course, supported!")
		e.HTML = []byte("<h1>你的验证码为1001!</h1>")
		err := e.SendWithTLS("smtp.163.com:587", smtp.PlainAuth("", "m17607466074@163.com", "HKZYGLMUUOCYKXIS", "smtp.163.com"),
			&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
		fmt.Println(err)
	}
}
