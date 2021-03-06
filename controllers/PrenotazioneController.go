package controllers

import (
	"Labooking/controllers/utils"
	"Labooking/models"
	"fmt"
	"github.com/beego/beego/v2/server/web"
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
			pc.Data["MaxTempi"] = pc.Data["MinTempi"].(int64) + 101
		} else {
			pc.Data["MaxTempi"] = it.Tempi/3600 + 1
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
	now, err := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
	if err != nil {
		pc.Ctx.WriteString("prenotazione: " + err.Error())
		return
	}
	if data.Before(now) {
		pc.Ctx.WriteString("prenotazione: data prenotazione scaduta")
		return
	}
	tipologiaTest := pc.GetString("tipologia-test")
	testDiagnostico.TipologiaTest = tipologiaTest

	switch fmt.Sprint(pc.GetSession("ruolo")) {
	case "privato":
		slotStr := pc.GetString("slot")
		slot, err := time.ParseInLocation("15:04", slotStr, time.Local)

		// verifica se lo slot ?? disponibile
		_ = testDiagnostico.SelezionaByDataStr(dataStr, slotStr)
		if testDiagnostico.IdTestDiagnostico != 0 {
			pc.Ctx.WriteString("prenotazione: slot gi?? prenotato!")
			return
		}
		testDiagnostico.DataEsecuzione = data.Add(time.Duration(slot.Hour())*time.Hour + time.Duration(slot.Minute())*time.Minute)
		nowStr := time.Now().Format("2006-01-02")
		now, err := time.ParseInLocation("2006-01-02", nowStr, time.Local)
		if err != nil {
			pc.Ctx.WriteString("prenotazione: " + err.Error())
			return
		}
		testDiagnostico.DataPrenotazione = now

		pagaOnline, err := pc.GetBool("metodo-pagamento")
		if err != nil {
			pc.Ctx.WriteString("prenotazione: " + err.Error())
			return
		}
		if pagaOnline {
			pagc := PagamentoController{}
			err = pagc.Paga(pc.Controller, testDiagnostico)
			if err != nil {
				pc.Ctx.WriteString("prenotazione: " + err.Error())
				return
			}
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
		qa.TestDiagnostico.IdTestDiagnostico = idTestDiagnostico
		_, err = qa.Aggiungi()
		if err != nil {
			pc.Ctx.WriteString("prenotazione: " + err.Error())
			return
		}
		testDiagnostico.IdTestDiagnostico = 0
		break
	case "organizzazione", "medico":
		slots := pc.GetStrings("slot")
		if len(slots) < 1 {
			pc.Ctx.WriteString("prenotazione: ?? necessario selezionare almeno un dipendente o paziente")
			return
		}

		pagaOnline, err := pc.GetBool("metodo-pagamento")
		if err != nil {
			pc.Ctx.WriteString("prenotazione: ?? necessario selezionare il metodo di pagamento desiderato")
			return
		}
		if pagaOnline {
			pagc := PagamentoController{}
			err = pagc.Paga(pc.Controller, testDiagnostico)
			if err != nil {
				pc.Ctx.WriteString("prenotazione: " + err.Error())
				return
			}
		}

		testDiagnostico.Stato = "prenotato"

		for _, slot := range slots {
			// verifica se lo slot ?? disponibile
			_ = testDiagnostico.SelezionaByDataStr(dataStr, slot)
			if testDiagnostico.IdTestDiagnostico != 0 {
				pc.Ctx.WriteString("prenotazione: slot gi?? prenotato!")
				return
			}

			slotTime, err := time.ParseInLocation("15:04", slot, time.Local)
			if err != nil {
				pc.Ctx.WriteString("prenotazione: " + err.Error())
				return
			}
			testDiagnostico.DataEsecuzione = data.Add(time.Duration(slotTime.Hour())*time.Hour + time.Duration(slotTime.Minute())*time.Minute)
			nowStr := time.Now().Format("2006-01-02")
			now, err := time.ParseInLocation("2006-01-02", nowStr, time.Local)
			if err != nil {
				pc.Ctx.WriteString("prenotazione: " + err.Error())
				return
			}
			testDiagnostico.DataPrenotazione = now

			testDiagnostico.Privato = new(models.Privato)
			idPrivatoStr := pc.GetString("privato-" + slot)
			idPrivato, err := strconv.Atoi(idPrivatoStr)
			if err != nil {
				pc.Ctx.WriteString("prenotazione: dipendente o paziente non valido")
				return
			}

			p := &models.Privato{IdPrivato: int64(idPrivato)}
			testDiagnostico.Privato = p
			idTestDiagnostico, err := testDiagnostico.Aggiungi()
			if err != nil {
				pc.Ctx.WriteString("prenotazione: " + err.Error())
				return
			}

			//creazione questionario
			file, fileHeaders, err := pc.GetFile("questionario-anamnesi-upload-" + slot)
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
			qa.TestDiagnostico.IdTestDiagnostico = idTestDiagnostico
			_, err = qa.Aggiungi()
			if err != nil {
				pc.Ctx.WriteString("prenotazione: " + err.Error())
				return
			}

			testDiagnostico.IdTestDiagnostico = 0
		}
		break
	default:
		pc.Ctx.WriteString("prenotazione: ruolo sconosciuto")
		return
	}

	pc.TplName = "dashboard/prenota/feedback.html"
}

type htmlSlot struct {
	Orario      string
	Disponibile bool
}

func costruisciSlot(allSlots, slotsPrenotati []*time.Time) []htmlSlot {
	var complexSlots []htmlSlot
	if len(slotsPrenotati) < 1 {
		for i, _ := range allSlots {
			complexSlot := htmlSlot{allSlots[i].Format("15:04"), true}
			complexSlots = append(complexSlots, complexSlot)
		}
		return complexSlots
	}

	for i, _ := range allSlots {
		flag := false
		for j, _ := range slotsPrenotati {
			if (*allSlots[i]).Local().Hour() == (*slotsPrenotati[j]).Hour() && (*allSlots[i]).Local().Minute() == (*slotsPrenotati[j]).Minute() {
				complexSlot := htmlSlot{allSlots[i].Format("15:04"), false}
				complexSlots = append(complexSlots, complexSlot)
				flag = true
				break
			}
		}
		if flag == false {
			complexSlot := htmlSlot{allSlots[i].Format("15:04"), true}
			complexSlots = append(complexSlots, complexSlot)
		}
	}

	return complexSlots
}

func (pc *PrenotazioneController) VisualizzaPrenotazioni() {
	td := new(models.TestDiagnostico)
	l := new(models.Laboratorio)
	email := pc.GetSession("email")
	l.Email = email.(string)
	err := l.Seleziona("email")
	if err != nil {
		pc.Ctx.WriteString("prenotazioni: " + err.Error())
		return
	}
	td.Laboratorio = l
	testDiagnostici, err := td.SelezionaTestAllByLab("id_test_diagnostico")
	if err != nil {
		return
	}
	utils.RenderLayout(&pc.Controller)
	pc.Data["Title"] = "Prenotazioni"
	pc.Data["TestDiagnostici"] = testDiagnostici
	pc.TplName = "dashboard/prenotazioni/prenotazioni_laboratorio.tpl"
}
