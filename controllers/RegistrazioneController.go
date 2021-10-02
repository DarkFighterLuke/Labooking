package controllers

import (
	"Labooking/models"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
)

type RegistrazioneController struct {
	web.Controller
}

func (rc *RegistrazioneController) Get() {
	rc.Data["FormPrivato"] = &models.Privato{}
	rc.Data["FormMedico"] = &models.Medico{}
	rc.Data["FormLaboratorio"] = &models.Laboratorio{}
	rc.TplName = "registrazione/registrazione.tpl"
}

func (rc *RegistrazioneController) Post() {
	switch rc.GetString("idForm") {
	case "privato":
		rc.registrazionePrivato()
		break
	case "medico":
		break
	case "laboratorio":
		break
	default:
		rc.Abort("400")
	}
}

func (rc *RegistrazioneController) registrazionePrivato() {
	p := models.Privato{}
	err := rc.ParseForm(&p)
	if err != nil {
		rc.Abort("400")
		return
	}
	valid := validation.Validation{}
	isValid, err := valid.Valid(&p)
	if err != nil {
		rc.Abort("500")
		return
	}
	if isValid {
		err = p.Aggiungi()
		if err != nil {
			rc.Abort("500")
			return
		}
	} else {
		for _, err := range valid.Errors {
			rc.Ctx.WriteString(err.Key + ": " + err.Message)
		}
	}
}
