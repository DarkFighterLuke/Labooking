package controllers

import (
	"Labooking/controllers/utils"
	"Labooking/models"
	"github.com/beego/beego/v2/server/web"
)

type PrivatoController struct {
	web.Controller
}

func (pc *PrivatoController) Get() {
	if pc.GetSession("ruolo") == "organizzazione" {
		o := new(models.Organizzazione)
		o.Email = pc.GetSession("email").(string)
		err := o.Seleziona("email")
		if err != nil {
			pc.Ctx.WriteString("visualizzazione dipendenti: " + err.Error())
			return
		}
		privati, err := o.GetDipendenti()
		if err != nil {
			pc.Ctx.WriteString("visualizzazione dipendenti: " + err.Error())
			return
		}
		pc.Data["Title"] = "I miei dipendenti"
		pc.Data["Privati"] = privati
	} else if pc.GetSession("ruolo") == "medico" {
		m := new(models.Medico)
		m.Email = pc.GetSession("email").(string)
		err := m.Seleziona("email")
		if err != nil {
			pc.Ctx.WriteString("visualizzazione pazienti: " + err.Error())
			return
		}
		privati, err := m.GetPazienti()
		if err != nil {
			pc.Ctx.WriteString("visualizzazione pazienti: " + err.Error())
			return
		}
		pc.Data["Title"] = "I miei pazienti"
		pc.Data["Privati"] = privati
	}
	utils.RenderLayout(&pc.Controller)
	pc.TplName = "dashboard/visualizzazionePrivati/visualizzazione_privati.tpl"
}

func (pc *PrivatoController) VisualizzaAggiunta() {
	utils.RenderLayout(&pc.Controller)
	if pc.GetSession("ruolo") == "organizzazione" {
		pc.Data["Title"] = "Aggiunta dipendente"
	} else if pc.GetSession("ruolo") == "medico" {
		pc.Data["Title"] = "Aggiunta paziente"
	}

	pc.Data["FormPrivato"] = &models.Privato{}
	pc.TplName = "registrazione/privato/registrazione_privato.tpl"
}
