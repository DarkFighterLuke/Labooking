package controllers

import (
	"Labooking/models"
	"github.com/beego/beego/v2/server/web"
)

type RegistrazioneController struct {
	web.Controller
}

func (rc *RegistrazioneController) Get() {
	rc.Data["FormPrivato"] = &models.Privato{}
	rc.Data["FormMedico"] = &models.Medico{}
	rc.TplName = "registrazione/registrazione.tpl"
}
