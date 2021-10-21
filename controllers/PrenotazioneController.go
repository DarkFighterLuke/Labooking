package controllers

import (
	"Labooking/controllers/utils"
	"Labooking/models"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"net/http"
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

		pc.Data["Data"] = time.Now().Format("2006-01-02")
		pc.Data["OraInizio"] = time.Now().Format("15:04")
		pc.Data["OraFine"] = time.Now().Add(time.Hour).Format("15:04")

		pc.LayoutSections["Head"] = "dashboard/prenota/head.html"
		pc.TplName = "dashboard/prenota/ricerca.tpl"
	} else if pc.GetString("action") == "prenotazione" {
		utils.RenderLayout(&pc.Controller)
		pc.Data["Title"] = "Prenota test"

		l := models.Laboratorio{}
		idLab, err := strconv.Atoi(pc.GetString("idLab"))
		if err != nil {
			pc.Ctx.WriteString("prenotazione: " + err.Error())
			return
		}
		pc.Data["IdLaboratorio"] = idLab

		l.IdLaboratorio = int64(idLab)
		err = l.Seleziona("id_laboratorio")
		if err != nil {
			pc.Ctx.WriteString("prenotazione: " + err.Error())
			return
		}

		var it models.InfoTest
		it.IdLaboratorio = new(models.Laboratorio)
		it.IdLaboratorio.IdLaboratorio = l.IdLaboratorio
		infoTest, err := it.SelezionaInfoTestByLabId()
		if err != nil {
			pc.Ctx.WriteString("prenotazione: " + err.Error())
			return
		}

		pc.Data["InfoTest"] = infoTest

		oraInizioStr := pc.GetString("inizio")
		oraFineStr := pc.GetString("fine")
		dataStr := pc.GetString("data")
		pc.Data["DataPrenotazione"] = dataStr
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

		if oraInizioStr == "" || oraFineStr == "" || dataStr == "" {
			pc.Ctx.WriteString("prenotazione: selezionare ora inizio o ora fine o data")
			return
		}

		isDisponibili, slots, slotsPrenotati, err := models.VerificaSlotDisponibili(l, oraInizioStr, oraFineStr, dataStr, numPersone)
		if err != nil {
			pc.Ctx.WriteString("prenotazione: " + err.Error())
			return
		}
		pc.Data["IsDisponibili"] = isDisponibili
		pc.Data["Slots"] = costruisciSlot(slots, slotsPrenotati)
		pc.Data["Iban"] = l.Iban
		ruolo := fmt.Sprint(pc.GetSession("ruolo"))
		switch ruolo {
		case "medico":
			m := new(models.Medico)
			m.Email = fmt.Sprint(pc.GetSession("email"))
			err := m.Seleziona("email")
			if err != nil {
				pc.Ctx.WriteString("prenotazione: " + err.Error())
				return
			}
			pazienti, err := m.GetPazienti()
			if err != nil {
				pc.Ctx.WriteString("prenotazione: " + err.Error())
				return
			}
			pc.Data["Privati"] = pazienti
			break
		case "organizzazione":
			org := new(models.Organizzazione)
			org.Email = fmt.Sprint(pc.GetSession("email"))
			err := org.Seleziona("email")
			if err != nil {
				pc.Ctx.WriteString("prenotazione: " + err.Error())
				return
			}
			dipendenti, err := org.GetDipendenti()
			if err != nil {
				pc.Ctx.WriteString("prenotazione: " + err.Error())
				return
			}
			pc.Data["Privati"] = dipendenti
			break
		}

		pc.TplName = "dashboard/prenota/prenota.tpl"
	}
}

func (pc *PrenotazioneController) Post() {
	testDiagnostico := new(models.TestDiagnostico)
	idLabStr := pc.GetString("id-laboratorio")
	idLab, err := strconv.Atoi(idLabStr)
	if err != nil {
		pc.Ctx.WriteString("prenotazione: " + err.Error())
		return
	}
	testDiagnostico.Laboratorio = new(models.Laboratorio)
	testDiagnostico.Laboratorio.IdLaboratorio = int64(idLab)
	dataStr := pc.GetString("data")
	data, err := time.ParseInLocation("2006-01-02", dataStr, time.Local)
	if err != nil {
		pc.Ctx.WriteString("prenotazione: " + err.Error())
		return
	}
	if data.Before(time.Now()) {
		pc.Ctx.WriteString("prenotazione: data prenotazione scaduta")
		return
	}

	switch fmt.Sprint(pc.GetSession("ruolo")) {
	case "privato":
		tipologiaTest := pc.GetString("tipologia-test")
		testDiagnostico.TipologiaTest = tipologiaTest
		slotStr := pc.GetString("slot")
		slot, err := time.ParseInLocation("15:04", slotStr, time.Local)

		testDiagnostico.DataEsecuzione = data.Add(time.Duration(slot.Hour())*time.Hour + time.Duration(slot.Minute())*time.Minute)

		pagaOnline, err := pc.GetBool("metodo-pagamento")
		if err != nil {
			pc.Ctx.WriteString("prenotazione: " + err.Error())
			return
		}
		if pagaOnline {
			fmt.Println(pc.GetString("numero-carta"),
				pc.GetString("scadenza"),
				pc.GetString("cvv"))
			testDiagnostico.Pagato = true
		} else {
			testDiagnostico.Pagato = false
		}

		p := new(models.Privato)
		p.Email = fmt.Sprint(pc.GetSession("email"))
		err = p.Seleziona("email")
		if err != nil {
			pc.Ctx.WriteString("prenotazione: " + err.Error())
			return
		}
		testDiagnostico.Privato = p

		testDiagnostico.Stato = "prenotato"

		idTestDiagnostico, err := testDiagnostico.Aggiungi()
		if err != nil {
			pc.Ctx.WriteString("prenotazione: " + err.Error())
			return
		}

		//creazione questionario
		file, fileHeaders, err := pc.GetFile("questionario-anamnesi-upload")
		if err != nil {
			pc.Ctx.WriteString("prenotazione: " + err.Error())
			return
		}
		fn := fileHeaders.Filename
		if fn[len(fn)-4:] != ".pdf" {
			pc.Ctx.WriteString("prenotazione: estensione file errata!")
			return
		}

		fileName, err := utils.SaveUploadedPdf(file, "pathquestionari")
		if err != nil {
			pc.Ctx.WriteString("prenotazione: " + err.Error())
			return
		}
		qa := new(models.QuestionarioAnamnesi)
		qa.Nome = fileName
		qa.TestDiagnostico = new(models.TestDiagnostico)
		qa.TestDiagnostico.IdTestDiagnostico = int(idTestDiagnostico)
		_, err = qa.Aggiungi()
		if err != nil {
			pc.Ctx.WriteString("prenotazione: " + err.Error())
			return
		}
		break
	case "laboratorio":
		break
	case "organizzazione":
		break
	}
	pc.Redirect("/dashboard/home", http.StatusFound)
}

type htmlSlot struct {
	Orario      string
	Disponibile bool
}

func costruisciSlot(allSlots, slotsPrenotati []*time.Time) []htmlSlot {
	var complexSlots []htmlSlot
	for i, _ := range allSlots {
		for j, _ := range slotsPrenotati {
			if (*allSlots[i]).Hour() == (*slotsPrenotati[j]).Hour() && (*allSlots[i]).Minute() == (*slotsPrenotati[j]).Minute() {
				complexSlot := htmlSlot{allSlots[i].Format("15:04"), false}
				complexSlots = append(complexSlots, complexSlot)
			} else {
				complexSlot := htmlSlot{allSlots[i].Format("15:04"), true}
				complexSlots = append(complexSlots, complexSlot)
			}
		}
	}

	return complexSlots
}
