package controllers

import (
	"Labooking/controllers/utils"
	"Labooking/models"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type LaboratorioController struct {
	web.Controller
}

func (lc *LaboratorioController) Get() {
	idLabStr := lc.GetString("idLab")
	idLab, err := strconv.Atoi(idLabStr)
	if idLabStr == "" || err != nil {
		lc.Abort("404")
		return
	}
	l := models.Laboratorio{IdLaboratorio: int64(idLab)}
	err = l.Seleziona()
	if err != nil {
		lc.Abort("500")
		return
	}

	var orariApertura []models.OrariApertura
	var orariChiusura []models.OrariApertura
	err = models.SelezionaOrariAperturaById(&orariApertura, idLab)
	err = models.SelezionaOrariChiusuraById(&orariChiusura, idLab)
	if err != nil {
		lc.Abort("500")
		return
	}

	type orariCustom struct {
		OrarioAperturaStr string
		OrarioChiusuraStr string
		Giorno            string
	}
	var orariStr []orariCustom

	for _, v := range orariApertura {
		orarioStr := v.Orario.Format("15:04")
		orariStr = append(orariStr, orariCustom{OrarioAperturaStr: orarioStr})
	}
	for i, v := range orariChiusura {
		orarioStr := v.Orario.Format("15:04")
		orariStr[i].OrarioChiusuraStr = orarioStr
		orariStr[i].Giorno = v.Giorno
	}

	var it []models.InfoTest
	err = models.SelezionaInfoTestByLabId(&it, idLab)
	if err != nil {
		lc.Abort("500")
		return
	}

	type infoTestCustom struct {
		TipologiaTest string
		Costo         float64
		TempiStr      string
	}
	var itCustom []infoTestCustom

	for _, v := range it {
		oreStr := strconv.Itoa(int(v.Tempi/3600)) + "h"
		minutiStr := ""
		if int(v.Tempi%3600) != 0 {
			minutiStr = strconv.Itoa(int(v.Tempi%3600)) + "m"
		}
		itCustom = append(itCustom, infoTestCustom{
			TipologiaTest: v.TipologiaTest,
			Costo:         v.Costo,
			TempiStr:      oreStr + " " + minutiStr,
		})
	}

	utils.RenderLayout(&lc.Controller)
	lc.Data["Title"] = l.Nome
	lc.Data["Orari"] = orariStr
	lc.Data["InfoTest"] = itCustom
	lc.TplName = "dashboard/laboratorio/laboratorio.tpl"
}
