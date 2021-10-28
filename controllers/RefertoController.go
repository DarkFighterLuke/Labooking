package controllers

import (
	"Labooking/controllers/utils"
	"Labooking/models"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type RefertoController struct {
	web.Controller
}

func (rc *RefertoController) Get() {
	if rc.GetSession("ruolo") == "laboratorio" {
		idRefertoStr := rc.GetString("idReferto")
		if idRefertoStr == "" {
			rc.Ctx.WriteString("referto: referto inesistente")
			return
		}
		r := new(models.Referto)
		idReferto, err := strconv.Atoi(idRefertoStr)
		if err != nil {
			rc.Ctx.WriteString("referto: " + err.Error())
			return
		}

		r.IdReferto = int64(idReferto)
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
			fmt.Println(err)
			return
		}
	}
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
		//TODO: invia mail privato
	}
	rc.Redirect("/dashboard/prenotazioni", http.StatusFound)
}
