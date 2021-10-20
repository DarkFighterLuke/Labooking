package controllers

import (
	"Labooking/controllers/utils"
	"Labooking/models"
	utils2 "Labooking/models/utils"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type PrenotazioneController struct {
	web.Controller
}

func (pc *PrenotazioneController) Get() {
	if pc.GetString("action") == "ricerca" {
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

		pc.Data["Ruolo"] = pc.GetSession("ruolo")

		pc.LayoutSections["Head"] = "dashboard/prenota/head.html"
		pc.TplName = "dashboard/prenota/ricerca.tpl"
	} else if pc.GetString("action") == "prenotazione" {
		utils.RenderLayout(&pc.Controller)
		pc.Data["Title"] = "Prenota test"
		pc.Data["Ruolo"] = pc.GetSession("ruolo")

		l := models.Laboratorio{}
		idLab, err := strconv.Atoi(pc.GetString("idLab"))
		if err != nil {
			pc.Ctx.WriteString("prenotazione: " + err.Error())
			return
		}
		l.IdLaboratorio = int64(idLab)
		err = l.Seleziona("id_laboratorio")
		if err != nil {
			pc.Ctx.WriteString("prenotazione: " + err.Error())
			return
		}
		pc.Data["Iban"] = l.Iban

		oraInizioStr := pc.GetString("inizio")
		oraFineStr := pc.GetString("fine")
		dataStr := pc.GetString("data")
		numPersoneStr := pc.GetString("persone")
		var numPersone int
		if numPersoneStr == "" {
			numPersone = 1
		} else {
			numPersone, err = strconv.Atoi(numPersoneStr)
			if err != nil {
				pc.Ctx.WriteString("prenotazione: " + err.Error())
				return
			}
		}

		isDisponibili, slots, slotsPrenotati, err := models.VerificaSlotDisponibili(l, oraInizioStr, oraFineStr, dataStr, numPersone)
		if err != nil {
			pc.Ctx.WriteString("prenotazione: " + err.Error())
			return
		}
		pc.Data["IsDisponibili"] = isDisponibili

		ruolo := fmt.Sprint(pc.GetSession("ruolo"))
		switch ruolo {
		case "medico":
			m := new(models.Medico)
			m.Email = fmt.Sprint(pc.GetSession("email"))
			err := m.Seleziona("email")
			if err != nil {
				pc.Ctx.WriteString("prenotazione: " + err.Error())
			}
			pazienti, err := m.GetPazienti()
			if err != nil {
				pc.Ctx.WriteString("prenotazione: " + err.Error())
			}
			pc.Data["Privati"] = pazienti
			pc.Data["Slots"] = costruisciSlot(slots, slotsPrenotati)
			break
		case "organizzazione":
			org := new(models.Organizzazione)
			org.Email = fmt.Sprint(pc.GetSession("email"))
			err := org.Seleziona("email")
			if err != nil {
				pc.Ctx.WriteString("prenotazione: " + err.Error())
			}
			dipendenti, err := org.GetDipendenti()
			if err != nil {
				pc.Ctx.WriteString("prenotazione: " + err.Error())
			}
			pc.Data["Privati"] = dipendenti
			pc.Data["Slots"] = costruisciSlot(slots, slotsPrenotati)
			break
		}

		pc.TplName = "dashboard/prenota/prenota.tpl"
	}
}

func (pc *PrenotazioneController) Post() {
	switch fmt.Sprint(pc.GetSession("ruolo")) {
	case "privato":
		file, _, err := pc.GetFile("questionario-anamnesi-upload")
		if err != nil {
			pc.Ctx.WriteString("prenotazione: " + err.Error())
		}
		err = os.MkdirAll("percorso", 0777)
		if err != nil {

		}
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			pc.Ctx.WriteString("prenotazione: " + err.Error())
		}

		fileName := utils2.RandStringRunes(32)
		err = ioutil.WriteFile(fileName+".pdf", fileBytes, 0655)
		if err != nil {
			pc.Ctx.WriteString("prenotazione: " + err.Error())
		}

		break
	case "laboratorio":
		break
	case "organizzazione":
		break
	}
}

type htmlSlot struct {
	Orario      string
	Disponibile bool
}

func costruisciSlot(allSlots, slotsPrenotati []*time.Time) []htmlSlot {
	var complexSlots []htmlSlot
	for i, _ := range allSlots {
		for j, _ := range slotsPrenotati {
			if *allSlots[i] == *slotsPrenotati[j] {
				complexSlot := htmlSlot{allSlots[i].Format("15:04"), false}
				complexSlots = append(complexSlots, complexSlot)
			}
		}
		complexSlot := htmlSlot{allSlots[i].Format("15:04"), true}
		complexSlots = append(complexSlots, complexSlot)
	}

	return complexSlots
}
