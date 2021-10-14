package controllers

import (
	"Labooking/controllers/utils"
	"github.com/beego/beego/v2/server/web"
)

type HomepageController struct {
	web.Controller
}

func (hc *HomepageController) Get() {
	utils.RenderLayout(&hc.Controller)
	hc.TplName = "dashboard/home/home.tpl"
}
