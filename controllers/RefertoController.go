package controllers

import (
	"Labooking/controllers/utils"
	"Labooking/models"
	"github.com/beego/beego/v2/server/web"
	"net/http"
	"strconv"
	"time"
)

type RefertoController struct {
	web.Controller
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
		r.IdReferto = int64(idTestDiagnostico)
		r.DataRilascio = time.Now()
		r.Risultato = esito
		_, err = r.Aggiungi()
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}
	}
	rc.Redirect("/dashboard/prenotazioni", http.StatusFound)
}
