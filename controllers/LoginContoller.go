package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	web.Controller
}

func (rc *LoginController) Get() {
	rc.TplName = "login/login.tpl"
}