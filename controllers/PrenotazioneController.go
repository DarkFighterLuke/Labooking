package controllers

import (
	"Labooking/controllers/utils"
	"Labooking/models"
	"github.com/beego/beego/v2/server/web"
)

type PrenotazioneController struct {
	web.Controller
}

func (pc *PrenotazioneController) Get() {
	utils.RenderLayout(&pc.Controller)
	pc.Data["Title"] = "Ricerca laboratori"
	it := &models.InfoTest{}
	err := it.SelezionaMinCosto()
	if err != nil {
		pc.Data["MinCosto"] = 0
	} else {
		pc.Data["MinCosto"] = it.Costo
	}

	err = it.SelezionaMaxCosto()
	if err != nil || it.Costo == pc.Data["MinCosto"].(float64) {
		pc.Data["MaxCosto"] = pc.Data["MinCosto"].(float64) + 100
	} else {
		pc.Data["MaxCosto"] = it.Costo
	}

	err = it.SelezionaMinTempi()
	if err != nil {
		pc.Data["MinTempi"] = 0
	} else {
		pc.Data["MinTempi"] = it.Tempi / 3600
	}

	err = it.SelezionaMaxTempi()
	if err != nil || (it.Tempi/3600) == pc.Data["MinTempi"].(int64) {
		pc.Data["MaxTempi"] = pc.Data["MinTempi"].(int64) + 100
	} else {
		pc.Data["MaxTempi"] = it.Tempi / 3600
	}

	pc.LayoutSections["Head"] = "dashboard/prenota/head.html"
	pc.TplName = "dashboard/prenota/ricerca.tpl"
}
