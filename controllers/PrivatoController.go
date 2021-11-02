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
	var privati []*models.Privato
	if pc.GetSession("ruolo") == "organizzazione" {
		o := new(models.Organizzazione)
		o.Email = pc.GetSession("email").(string)
		err := o.Seleziona("email")
		if err != nil {
			pc.Ctx.WriteString("visualizzazione dipendenti: " + err.Error())
			return
		}
		privati, err = o.GetDipendenti()
		if err != nil {
			pc.Ctx.WriteString("visualizzazione dipendenti: " + err.Error())
			return
		}

		pc.Data["Title"] = "I miei dipendenti"
	} else if pc.GetSession("ruolo") == "medico" {

	}
	utils.RenderLayout(&pc.Controller)
	pc.Data["Privati"] = privati
	pc.TplName = "dashboard/visualizzazionePrivati/visualizzazione_privati.tpl"
}

func (pc *PrivatoController) VisualizzaAggiunta() {
	utils.RenderLayout(&pc.Controller)
	pc.Data["Title"] = "Aggiunta dipendente"
	pc.Data["FormPrivato"] = &models.Privato{}
	pc.TplName = "registrazione/privato/registrazione_privato.tpl"
}
