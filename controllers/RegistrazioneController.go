package controllers

import (
	"Labooking/models"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
	"github.com/pkg/errors"
	"net/http"
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
		err := rc.registrazionePrivato()
		if err != nil {
			rc.Ctx.WriteString(err.Error())
			return
		}
		rc.Redirect("/dashboard?page=home", http.StatusFound)
		break
	case "medico":
		err := rc.registrazioneMedico()
		if err != nil {
			rc.Ctx.WriteString(err.Error())
			return
		}
		rc.Redirect("/dashboard?page=home", http.StatusFound)
		break
	case "laboratorio":
		err := rc.registrazioneLaboratorio()
		if err != nil {
			rc.Ctx.WriteString(err.Error())
			return
		}
		rc.Redirect("/dashboard?page=home", http.StatusFound)
		break
	case "organizzazione":
		//TODO: implementa logica registrazione organizzazione
		break
	default:
		rc.Abort("400")
	}
}

func (rc *RegistrazioneController) registrazionePrivato() error {
	p := models.Privato{}
	err := rc.ParseForm(&p)
	if err != nil {
		return err
	}
	p.Prefisso = rc.GetString("Prefisso")
	err = rc.validateAndInsert(&p)
	if err != nil {
		return err
	}
	return nil
}

func (rc *RegistrazioneController) registrazioneMedico() error {
	m := models.Medico{}
	err := rc.ParseForm(&m)
	if err != nil {
		return err
	}
	m.Prefisso = rc.GetString("Prefisso")
	err = rc.validateAndInsert(&m)
	if err != nil {
		return err
	}
	return nil
}

func (rc *RegistrazioneController) registrazioneLaboratorio() error {
	l := models.Laboratorio{}
	err := rc.ParseForm(&l)
	if err != nil {
		return err
	}
	l.Prefisso = rc.GetString("Prefisso")
	err = rc.validateAndInsert(&l)
	if err != nil {
		return err
	}
	return nil
}

func (rc *RegistrazioneController) validateAndInsert(user models.ReadWriteDB) error {
	valid := validation.Validation{}
	isValid, err := valid.Valid(user)
	if err != nil {
		return err
	}
	if isValid {
		err = user.Aggiungi()
		if err != nil {
			return err
		}
	} else {
		var errori string
		for _, err := range valid.Errors {
			errori = errori + err.Key + ":" + err.Message + "\n"
		}
		return errors.Errorf(errori)
	}
	return nil
}
