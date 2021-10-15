package utils

import (
	"github.com/beego/beego/v2/server/web"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"time"
)

func IsAdult(dataNascita time.Time) bool {
	if time.Now().Sub(dataNascita).Hours() < 157680 {
		return false
	} else {
		return true
	}
}

func IsPswValid(psw string) bool {
	entropy, err := web.AppConfig.Float("pswentropy")
	if err != nil {
		return false
	}
	if err := passwordvalidator.Validate(psw, entropy); err != nil {
		return false
	} else {
		return true
	}
}
