package api

import (
	"Labooking/models"
	"github.com/beego/beego/v2/server/web"
	"time"
)

type AziendaSanitaria struct {
	web.Controller
}

func (as *AziendaSanitaria) PrelevaReferti() {
	testDiagnostici, err := models.GetRefertiByTempoLuogo()
	if err != nil {
		as.Ctx.WriteString("preleva referti: " + err.Error())
		return
	}

	type servedJson struct {
		DataEsecuzione       time.Time
		TipologiaTest        string
		NomeLaboratorio      string
		IndirizzoLaboratorio string
	}
	var jsons []*servedJson

	for _, v := range testDiagnostici {
		json := &servedJson{v.DataEsecuzione, v.TipologiaTest, v.Laboratorio.Nome, v.Laboratorio.Indirizzo}
		jsons = append(jsons, json)
	}

	as.Data["json"] = jsons
	err = as.ServeJSON()
	if err != nil {
		as.Ctx.WriteString("preleva referti: " + err.Error())
		return
	}
}
