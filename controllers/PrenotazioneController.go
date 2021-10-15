package controllers

import (
	"Labooking/controllers/utils"
	"github.com/beego/beego/v2/server/web"
)

type PrenotazioneController struct {
	web.Controller
}

func (pc *PrenotazioneController) Get() {
	utils.RenderLayout(&pc.Controller)
	pc.Data["Title"] = "Ricerca laboratori"
	pc.LayoutSections["Head"] = "dashboard/prenota/head.html"
	pc.TplName = "dashboard/prenota/ricerca.tpl"
}
