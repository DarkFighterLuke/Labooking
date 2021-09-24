package controllers

import "github.com/astaxie/beego"

type HomepageController struct {
	beego.Controller
}

func (c *HomepageController) Get() {
	c.Data["Title"] = "Attività recenti"
	c.Data["TabName"] = "home"
	c.TplName = "dashboard/page_content_wrapper.tpl"
}
