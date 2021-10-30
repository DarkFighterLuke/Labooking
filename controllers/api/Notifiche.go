package api

import (
	"Labooking/models"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type Notifiche struct {
	web.Controller
}

func (n *Notifiche) UpdateNotificheLette() {
	for i := 0; ; i++ {
		idTestDiagnosticoStr := n.GetString("notifica-" + strconv.Itoa(i))
		if idTestDiagnosticoStr == "" {
			n.Ctx.ResponseWriter.WriteHeader(200)
		}
		idTestDiagnostico, err := strconv.Atoi(idTestDiagnosticoStr)
		if err != nil {
			continue
		}
		td := new(models.TestDiagnostico)
		td.IdTestDiagnostico = int64(idTestDiagnostico)
		err = td.Seleziona("id_test_diagnostico")
		if err != nil {
			continue
		}
		td.Stato = "visualizzato"
		_ = td.Modifica()
	}
}
