package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"net/smtp"
	"strings"
)

func InviaMail(msg string, recivers []string) error {
	from, err := web.AppConfig.String("email")
	if err != nil {
		return err
	}
	pass, err := web.AppConfig.String("passwordemail")
	if err != nil {
		return err
	}

	var to string
	if len(recivers) == 1 {
		to = recivers[len(recivers)-1]
	} else {
		to = strings.Join(recivers, ",")
	}

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
