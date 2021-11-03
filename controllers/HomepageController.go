package controllers

import (
	"Labooking/controllers/utils"
	"Labooking/models"
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

type HomepageController struct {
	web.Controller
}

func (hc *HomepageController) Get() {
	utils.RenderLayout(&hc.Controller)

	td := new(models.TestDiagnostico)
	utente := fmt.Sprint(hc.GetSession("ruolo"))
	switch utente {
	case "privato":
		p := new(models.Privato)

		p.Email = fmt.Sprint(hc.GetSession("email"))
		err := p.Seleziona("email")
		if err != nil {
			hc.Ctx.WriteString("homepage: " + err.Error())
			return
		}
		td.Privato = p
		break
	case "laboratorio":
		l := new(models.Laboratorio)

		l.Email = fmt.Sprint(hc.GetSession("email"))
		err := l.Seleziona("email")
		if err != nil {
			hc.Ctx.WriteString("homepage: " + err.Error())
			return
		}
		td.Laboratorio = l
		break
	case "medico":
		m := new(models.Medico)

		m.Email = fmt.Sprint(hc.GetSession("email"))
		err := m.Seleziona("email")
		if err != nil {
			hc.Ctx.WriteString("homepage: " + err.Error())
			return
		}
		td.Privato.Medico = m
		break
	case "organizzazione":
		o := new(models.Organizzazione)

		o.Email = fmt.Sprint(hc.GetSession("email"))
		err := o.Seleziona("email")
		if err != nil {
			hc.Ctx.WriteString("homepage: " + err.Error())
			return
		}
		td.Privato.Organizzazione = o
		break
	}

	tests, err := td.SelezionaLastUpdate(10, utente)
	if err != nil {
		hc.Ctx.WriteString("homepage: " + err.Error())
		return
	}
	hc.Data["TestDiagnostici"] = tests
	hc.Data["Title"] = "Attività recenti"
	hc.TplName = "dashboard/home/home.tpl"
}
