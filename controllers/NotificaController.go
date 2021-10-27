package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"net/smtp"
)

func InviaMail(msg, reciver string) error {
	from, err := web.AppConfig.String("email")
	if err != nil {
		return err
	}
	pass, err := web.AppConfig.String("passwordemail")
	if err != nil {
		return err
	}
	to := reciver

	msg = "From: " + from + "\n" +
		"To: " + to + "\n" + msg

	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		return err
	}
	return nil
}
