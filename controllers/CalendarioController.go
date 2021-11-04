package controllers

import (
	"Labooking/controllers/utils"
	"Labooking/models"
	"github.com/beego/beego/v2/server/web"
	"strconv"
	"time"
)

type CalendarioController struct {
	web.Controller
}

func (cc *CalendarioController) Get() {
	p := new(models.Privato)
	p.Email = cc.GetSession("email").(string)
	err := p.Seleziona("email")
	if err != nil {
		cc.Ctx.WriteString("calendario: " + err.Error())
		return
	}
	td := new(models.TestDiagnostico)
	td.Privato = p
	testDiagnostici, err := td.SelezionaTestAllByPriv()
	if err != nil {
		cc.Ctx.WriteString("calendario: " + err.Error())
		return
	}

	var eventi []models.CalendarEvent
	i := 0
	for _, v := range testDiagnostici {
		eventoPrenotazione := models.CalendarEvent{
			Id:    i,
			Title: "Prenotato Test " + v.TipologiaTest,
			Url:   "",
			Class: "event-info",
			Start: v.DataPrenotazione.Add(1 * time.Second).UnixMilli(),
			End:   v.DataPrenotazione.Add(1 * time.Second).UnixMilli(),
		}
		eventi = append(eventi, eventoPrenotazione)
		i++

		eventoEsecuzione := models.CalendarEvent{
			Id:    i,
			Title: "Esecuzione Test " + v.TipologiaTest + " pianificata per questa data!",
			Url:   "",
			Class: "event-warning",
			Start: v.DataEsecuzione.UnixMilli(),
			End:   v.DataEsecuzione.UnixMilli(),
		}
		eventi = append(eventi, eventoEsecuzione)
		i++

		if v.Referto != nil {
			eventoReferto := models.CalendarEvent{
				Id:    i,
				Title: "Referto Test " + v.TipologiaTest + " pubblicato in questa data!",
				Url:   "/dashboard/referti/download?idReferto=" + strconv.Itoa(int(v.Referto.IdReferto)),
				Class: "event-success",
				Start: v.DataEsecuzione.UnixMilli(),
				End:   v.DataEsecuzione.UnixMilli(),
			}
			eventi = append(eventi, eventoReferto)
			i++
		}
	}

	utils.RenderLayout(&cc.Controller)
	cc.Data["Title"] = "Calendario"
	cc.LayoutSections["Head"] = "dashboard/calendario/head.tpl"
	cc.Data["Eventi"] = eventi
	cc.TplName = "dashboard/calendario/calendario.tpl"
}
