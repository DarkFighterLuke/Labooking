package controllers

import (
	"Labooking/controllers/utils"
	"Labooking/models"
	"github.com/beego/beego/v2/server/web"
)

type DipendentiController struct {
	web.Controller
}

func (dc *DipendentiController) Get() {
	o := new(models.Organizzazione)
	o.Email = dc.GetSession("email").(string)
	err := o.Seleziona("email")
	if err != nil {
		dc.Ctx.WriteString("visualizzazione dipendenti: " + err.Error())
		return
	}
	dipendenti, err := o.GetDipendenti()
	if err != nil {
		dc.Ctx.WriteString("visualizzazione dipendenti: " + err.Error())
		return
	}
	utils.RenderLayout(&dc.Controller)
	dc.Data["Title"] = "I miei dipendenti"
	dc.Data["Privati"] = dipendenti
	dc.TplName = "dashboard/visualizzazionePrivati/visualizzazione_privati.tpl"
}

func (dc *DipendentiController) VisualizzaAggiunta() {
	utils.RenderLayout(&dc.Controller)
	dc.Data["Title"] = "Aggiunta dipendente"
	dc.Data["FormPrivato"] = &models.Privato{}
	dc.TplName = "registrazione/privato/registrazione_privato.tpl"
}
