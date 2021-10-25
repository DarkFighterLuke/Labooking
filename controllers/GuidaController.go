package controllers

import (
	"Labooking/controllers/utils"
	"github.com/beego/beego/v2/server/web"
)

type GuidaController struct {
	web.Controller
}

func (gc *GuidaController) Get() {
	utils.RenderLayout(&gc.Controller)
	gc.Data["Title"] = "Guida ai test"
	gc.TplName = "dashboard/guida/guida.tpl"
}
