package controllers

import "github.com/beego/beego/v2/server/web"

type LandingPageController struct {
	web.Controller
}

func (c *LandingPageController) Get() {
	c.TplName = "landingPage/index.tpl"
}
