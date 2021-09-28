package controllers

import "github.com/astaxie/beego"

type LandingPageController struct {
	beego.Controller
}

func (c *LandingPageController) Get() {
	c.TplName = "landingPage/index.html"
}
