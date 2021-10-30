package api

import (
	"Labooking/models"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type Notifiche struct {
	web.Controller
}

func (nc *Notifiche) UpdateNotificheLette() {
	for i := 0; ; i++ {
		idTestDiagnosticoStr := nc.GetString("notifica-" + strconv.Itoa(i))
		if idTestDiagnosticoStr == "" {
			return
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
