package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"net/http"
)

type LogoutController struct {
	web.Controller
}

func (lo *LogoutController) Get() {
	err := lo.DelSession("email")
	err = lo.DelSession("ruolo")
	if err != nil {
		lo.Ctx.WriteString("logout: " + err.Error())
	}
	lo.Redirect("/", http.StatusFound)
}
