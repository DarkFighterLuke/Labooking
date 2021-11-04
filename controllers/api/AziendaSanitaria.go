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
		Risultato            string
		DataRilascio         time.Time
	}
	var jsons []*servedJson

	for _, v := range testDiagnostici {
		json := &servedJson{v.DataEsecuzione, v.TipologiaTest, v.Laboratorio.Nome, v.Laboratorio.Indirizzo, v.Referto.Risultato, v.Referto.DataRilascio}
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

func (as *AziendaSanitaria) PrelevaStatisticheLaboratori() {
	laboratori, err := models.SelezionaAllLaboratori()
	if err != nil {
		as.Ctx.WriteString("preleva statistiche laboratori: " + err.Error())
		return
	}

	type statisticaLaboratorio struct {
		PartitaIva                      string
		Nome                            string
		NumeroTestDiagnosticiEffettuati int
		NumeroPositivi                  int
		NumeroNegativi                  int
		NumeroNulli                     int
	}

	var statisticheLaboratori []*statisticaLaboratorio
	for _, v := range laboratori {
		td := models.TestDiagnostico{Laboratorio: v}

		testsByLab, err := td.SelezionaTestAllByLab("id_laboratorio")
		if err != nil {
			as.Ctx.WriteString("preleva statistiche laboratori: " + err.Error())
			return
		}

		var statisticaLaboratorio statisticaLaboratorio
		statisticaLaboratorio.PartitaIva = v.PartitaIva
		statisticaLaboratorio.Nome = v.Nome
		for _, k := range testsByLab {
			if k.Referto != nil {
				switch k.Referto.Risultato {
				case "positivo":
					statisticaLaboratorio.NumeroTestDiagnosticiEffettuati++
					statisticaLaboratorio.NumeroPositivi++
					break
				case "negativo":
					statisticaLaboratorio.NumeroTestDiagnosticiEffettuati++
					statisticaLaboratorio.NumeroNegativi++
					break
				case "nullo":
					statisticaLaboratorio.NumeroTestDiagnosticiEffettuati++
					statisticaLaboratorio.NumeroNulli++
					break
				default:
				}
			}
		}
		statisticheLaboratori = append(statisticheLaboratori, &statisticaLaboratorio)
	}

	as.Data["json"] = statisticheLaboratori
	err = as.ServeJSON()
	if err != nil {
		as.Ctx.WriteString("preleva statistiche laboratori: " + err.Error())
		return
	}
}
