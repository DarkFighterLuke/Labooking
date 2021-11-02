package controllers

import (
	"Labooking/controllers/utils"
	"Labooking/models"
	"github.com/beego/beego/v2/server/web"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type RefertoController struct {
	web.Controller
}

func (rc *RefertoController) Get() {
	var testDiagnostici []*models.TestDiagnostico
	email := rc.GetSession("email")
	if rc.GetSession("ruolo") == "privato" {
		td := new(models.TestDiagnostico)
		p := new(models.Privato)
		p.Email = email.(string)
		err := p.Seleziona("email")
		if err != nil {
			rc.Ctx.WriteString("referti: " + err.Error())
			return
		}
		td.Privato = p
		testDiagnostici, err = td.SelezionaTestAllByPriv()
		if err != nil {
			return
		}
	} else if rc.GetSession("ruolo") == "medico" {
		m := new(models.Medico)
		m.Email = email.(string)
		err := m.Seleziona("email")
		if err != nil {
			rc.Ctx.WriteString("referti: " + err.Error())
			return
		}
		testDiagnostici, err = models.SelezionaTestAllByMed(m.IdMedico)
		if err != nil {
			return
		}
	} else if rc.GetSession("ruolo") == "organizzazione" {
		o := new(models.Organizzazione)
		o.Email = email.(string)
		err := o.Seleziona("email")
		if err != nil {
			rc.Ctx.WriteString("referti: " + err.Error())
			return
		}
		testDiagnostici, err = models.SelezionaTestAllByOrg(o.IdOrganizzazione)
		if err != nil {
			return
		}
	}

	utils.RenderLayout(&rc.Controller)
	rc.Data["Title"] = "Referti"
	rc.Data["TestDiagnostici"] = testDiagnostici
	rc.TplName = "dashboard/prenotazioni/prenotazioni_laboratorio.tpl"
}

func (rc *RefertoController) Post() {
	for i := 0; ; i++ {
		idTDStr := rc.GetString("test-diagnostico-" + strconv.Itoa(i))
		if idTDStr == "" {
			break
		}
		idTestDiagnostico, err := strconv.Atoi(idTDStr)
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}
		referto, _, err := rc.GetFile("referto-upload-" + strconv.Itoa(i))
		if err == http.ErrMissingFile {
			continue
		} else if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}
		filename, err := utils.SaveUploadedPdf(referto, "pathreferti")
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}
		esito := rc.GetString("esito-" + strconv.Itoa(i))
		if esito != "nullo" && esito != "negativo" && esito != "positivo" {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}
		r := new(models.Referto)
		r.Nome = filename
		r.DataRilascio = time.Now()
		r.Risultato = esito
		idReferto, err := r.Aggiungi()
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}

		testDiagnostico := new(models.TestDiagnostico)
		testDiagnostico.IdTestDiagnostico = int64(idTestDiagnostico)
		err = testDiagnostico.Seleziona("id_test_diagnostico")
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}
		testDiagnostico.Stato = "notificato"
		testDiagnostico.Referto = &models.Referto{IdReferto: idReferto}
		err = testDiagnostico.Modifica()
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}

		//componi mail privato
		err = testDiagnostico.Seleziona("id_test_diagnostico")
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}

		privato := new(models.Privato)
		privato.IdPrivato = testDiagnostico.Privato.IdPrivato
		err = privato.Seleziona("id_privato")
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}

		msgPrivato, err := componiMsgPrivato(int(r.IdReferto), r.Risultato)
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}

		//componi mail privato
		msgMedico, err := componiMsgMedico(int(r.IdReferto), int(privato.IdPrivato), r.Risultato)
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}

		medico := new(models.Medico)
		medico.IdMedico = privato.Medico.IdMedico
		err = medico.Seleziona("id_medico")
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}

		//componi mail organizzazione
		organizzazione := new(models.Organizzazione)
		organizzazione.IdOrganizzazione = privato.Organizzazione.IdOrganizzazione
		err = organizzazione.Seleziona()
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}
		testDiagnostico.Privato.Organizzazione = new(models.Organizzazione)
		testDiagnostico.Privato.Organizzazione.IdOrganizzazione = organizzazione.IdOrganizzazione
		ok, err := testDiagnostico.CheckInviaMailiOrganizzazione()
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}

		var msgOrganizzazione string
		if ok {
			msgOrganizzazione, err = componiMsgOrganizzazione(testDiagnostico.DataPrenotazione)
			if err != nil {
				rc.Ctx.WriteString("referto: " + err.Error())
				return
			}
		}

		//invio mail
		err = InviaMail(msgPrivato, []string{privato.Email})
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}

		err = InviaMail(msgMedico, []string{medico.Email})
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}

		if ok {
			err = InviaMail(msgOrganizzazione, []string{organizzazione.Email})
			if err != nil {
				rc.Ctx.WriteString("referto: " + err.Error())
				return
			}
		}

	}
	rc.Redirect("/dashboard/prenotazioni", http.StatusFound)
}

func componiMsgPrivato(idReferto int, esito string) (string, error) {
	websitelink, err := web.AppConfig.String("websitelink")
	if err != nil {
		return "", err
	}
	idRefertoStr := strconv.Itoa(idReferto)
	link := websitelink + "dashboard/referti?idReferto=" + idRefertoStr

	msg := "Subject: Il tuo referto è disponibile\n\n" +
		"Il tuo referto è pronto e l'esito è: " + esito + "\n" +
		"Clicca qui per visualizzarlo: " + link + "\n"

	return msg, nil
}

func componiMsgMedico(idReferto int, paziente int, esito string) (string, error) {
	websitelink, err := web.AppConfig.String("websitelink")
	if err != nil {
		return "", err
	}
	idRefertoStr := strconv.Itoa(idReferto)
	link := websitelink + "dashboard/referti?idReferto=" + idRefertoStr
	pazienteStr := strconv.Itoa(paziente)
	msg := "Subject: Referto paziente " + pazienteStr + " è disponibile\n\n" +
		"Il referto del paziente " + pazienteStr + " è pronto e l'esito è: " + esito + "\n" +
		"Clicca qui per visualizzarlo: " + link + "\n"

	return msg, nil
}

func componiMsgOrganizzazione(dataPrenotazione time.Time) (string, error) {
	websitelink, err := web.AppConfig.String("websitelink")
	if err != nil {
		return "", err
	}
	link := websitelink + "login"
	dataStr := dataPrenotazione.Format("2006-01-02")
	msg := "Subject: Referti prenotato il " + dataStr + " disponibili\n\n" +
		"I referti dei tuoi dipendenti relativi alla prenotazione effettuata in data " + dataStr + " sono pronti\n" +
		"Visita il nostro sito per visualizzarli: " + link + "\n"

	return msg, nil
}

func (rc *RefertoController) DownloadReferto() {
	idTestDiagnosticoStr := rc.GetString("idReferto")
	if idTestDiagnosticoStr == "" {
		rc.Ctx.WriteString("referto: referto inesistente")
		return
	}

	idTestDiagnostico, err := strconv.Atoi(idTestDiagnosticoStr)
	if err != nil {
		rc.Ctx.WriteString("referto: " + err.Error())
		return
	}

	td := new(models.TestDiagnostico)
	td.Referto = &models.Referto{IdReferto: int64(idTestDiagnostico)}
	err = td.Seleziona("id_referto")
	if err != nil {
		rc.Ctx.WriteString("referto: referto inesistente")
		return
	}

	ruolo := rc.GetSession("ruolo").(string)
	email := rc.GetSession("email").(string)
	if ruolo == "laboratorio" {
		l := new(models.Laboratorio)
		l.Email = email
		err = l.Seleziona("email")
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}
		if l.IdLaboratorio != td.Laboratorio.IdLaboratorio {
			rc.Ctx.WriteString("referto: non sei autorizzato a visualizzare questo referto")
			return
		}
	} else if ruolo == "privato" {
		p := new(models.Privato)
		p.Email = email
		err = p.Seleziona("email")
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}
		if p.IdPrivato != td.Privato.IdPrivato {
			rc.Ctx.WriteString("referto: non sei autorizzato a visualizzare questo referto")
			return
		}
	} else if ruolo == "medico" {
		m := new(models.Medico)
		m.Email = email
		err = m.Seleziona("email")
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}
		err = td.Privato.Seleziona("id_privato")
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}
		if td.Privato.Medico.IdMedico != m.IdMedico {
			rc.Ctx.WriteString("referto: non sei autorizzato a visualizzare questo referto")
			return
		}
	} else if ruolo == "organizzazione" {
		o := new(models.Organizzazione)
		o.Email = email
		err = o.Seleziona("email")
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}
		err = td.Privato.Seleziona("id_privato")
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}
		if td.Privato.Organizzazione.IdOrganizzazione != o.IdOrganizzazione {
			rc.Ctx.WriteString("referto: non sei autorizzato a visualizzare questo referto")
			return
		}
	}

	r := new(models.Referto)
	r.IdReferto = td.Referto.IdReferto
	err = r.Seleziona("id_referto")
	if err != nil {
		rc.Ctx.WriteString("referto: referto inesistente")
		return
	}

	pathReferti, err := web.AppConfig.String("pathreferti")
	if err != nil {
		rc.Ctx.WriteString("referto: " + err.Error())
		return
	}
	refertoBytes, err := ioutil.ReadFile(pathReferti + r.Nome + ".pdf")
	if err != nil {
		rc.Ctx.WriteString("referto: " + err.Error())
		return
	}

	rc.Ctx.ResponseWriter.Header().Set("Content-Type", "application/pdf")
	_, err = rc.Ctx.ResponseWriter.Write(refertoBytes)
	if err != nil {
		log.Println(err)
		return
	}
}
