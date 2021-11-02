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

func (as *AziendaSanitaria) PrelevaUtentiPositivi() {
	privati, err := models.GetPrivatiPositivi()
	if err != nil {
		as.Ctx.WriteString("preleva utenti positivi: " + err.Error())
		return
	}

	type privatoMinified struct {
		Nome                   string
		Cognome                string
		CodiceFiscale          string
		NumeroTesseraSanitaria string
		Indirizzo              string
		Prefisso               string
		Telefono               string
		Email                  string
	}

	var privatiMinified []*privatoMinified
	for _, v := range privati {
		privatiMinified = append(privatiMinified, &privatoMinified{v.Nome, v.Cognome, v.CodiceFiscale, v.NumeroTesseraSanitaria, v.Indirizzo, v.Prefisso, v.Telefono, v.Email})
	}

	as.Data["json"] = privatiMinified
	err = as.ServeJSON()
	if err != nil {
		as.Ctx.WriteString("preleva utenti positivi: " + err.Error())
		return
	}
}
