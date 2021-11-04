package controllers

import (
	"Labooking/controllers/utils"
	"Labooking/models"
	"github.com/beego/beego/v2/server/web"
	"strconv"
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

func (pc *PrivatoController) Eliminazione() {
	idPrivatiStr := pc.GetStrings("id-privato")
	for _, idPrivatoStr := range idPrivatiStr {
		if idPrivatoStr == "" {
			break
		}
		idPrivato, err := strconv.Atoi(idPrivatoStr)
		if err != nil {
			pc.Ctx.WriteString("eliminazione pazienti: " + err.Error())
			return
		}
		p := new(models.Privato)
		p.IdPrivato = int64(idPrivato)
		if pc.GetSession("ruolo") == "organizzazione" {
			o := new(models.Organizzazione)
			o.Email = pc.GetSession("email").(string)
			err = o.Seleziona("email")
			if err != nil {
				pc.Ctx.WriteString("eliminazione dipendenti: " + err.Error())
				return
			}
			p.Organizzazione = o
			err = p.Seleziona("id_privato", "organizzazione")
			if err != nil {
				pc.Ctx.WriteString("eliminazione dipendenti: " + err.Error())
				return
			}

			p.Organizzazione = nil
			err = p.Modifica()
			if err != nil {
				pc.Ctx.WriteString("eliminazione dipendenti: " + err.Error())
				return
			}
		} else if pc.GetSession("ruolo") == "medico" {
			m := new(models.Medico)
			m.Email = pc.GetSession("email").(string)
			err = m.Seleziona("email")
			if err != nil {
				pc.Ctx.WriteString("eliminazione pazienti: " + err.Error())
				return
			}
			p.Medico = m
			err = p.Seleziona("id_privato", "medico")
			if err != nil {
				pc.Ctx.WriteString("eliminazione pazienti: " + err.Error())
				return
			}

			p.Medico = nil
			err = p.Modifica()
			if err != nil {
				pc.Ctx.WriteString("eliminazione dipendenti: " + err.Error())
				return
			}
		}
	}

	pc.Ctx.ResponseWriter.WriteHeader(200)
}
