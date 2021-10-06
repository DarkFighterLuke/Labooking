package controllers

import (
	"Labooking/models"
	"fmt"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"time"
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
	err = rc.validate(&p)
	if err != nil {
		return err
	}
	_, err = p.Aggiungi()
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
	err = rc.validate(&m)
	if err != nil {
		return err
	}
	_, err = m.Aggiungi()
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
	orari, err := rc.parseAndValidateOrariApertura()
	if err != nil {
		return err
	}

	err = rc.validate(&l)
	if err != nil {
		return err
	}
	id, err := l.Aggiungi()
	if err != nil {
		return err
	}

	for _, o := range orari {
		o.IdLaboratorio = &models.Laboratorio{IdLaboratorio: id}
		_, err = o.Aggiungi()
		if err != nil {
			continue
		}
	}
	return nil
}

func (rc *RegistrazioneController) parseAndValidateOrariApertura() ([]models.OrariApertura, error) {
	var orari []models.OrariApertura
	i := 1
	var orario models.OrariApertura
	var orarioApertura = " "
	var orarioChiusura = " "
	var giorno = " "
	for orarioApertura != "" || orarioChiusura != "" || giorno != "" {
		orarioApertura = rc.GetString("orario-apertura-" + strconv.Itoa(i))
		orarioChiusura = rc.GetString("orario-chiusura-" + strconv.Itoa(i))
		giorno = rc.GetString("giorno-" + strconv.Itoa(i))
		if orarioApertura == "" || orarioChiusura == "" || giorno == "" {
			break
		} else {
			oa, err := time.Parse("03:04", orarioApertura)
			if err != nil {
				return nil, err
			}
			oa = oa.AddDate(1, 0, 0)
			oc, err := time.Parse("03:04", orarioChiusura)
			if err != nil {
				return nil, err
			}
			oc = oc.AddDate(1, 0, 0)
			if !oc.After(oa) {
				return nil, fmt.Errorf("%s", "orario apertura successivo a orario di chiusura")
			}
			// Controlla che gli intervalli non si accavallino
			for j := 0; j < len(orari); j += 2 {
				oatemp := orari[j].Orario
				octemp := orari[j+1].Orario
				gtemp := orari[j].Giorno
				if oa.After(oatemp) && oa.Before(octemp) && gtemp == giorno || oc.After(oatemp) && oc.Before(octemp) && gtemp == giorno {
					return nil, fmt.Errorf("%s", "intervallo contenuto in un altro intervallo già presente")
				}
			}
			orario.Orario = oa
			orario.Stato = true
			orario.Giorno = giorno
			orari = append(orari, orario)
			orario2 := orario
			orario2.Orario = oc
			orario2.Stato = false
			orari = append(orari, orario2)
			i++
		}
	}
	return orari, nil
}

func (rc *RegistrazioneController) validate(user models.ReadWriteDB) error {
	valid := validation.Validation{}
	isValid, err := valid.Valid(user)
	if err != nil {
		return err
	}
	if !isValid {
		var errori string
		for _, err := range valid.Errors {
			errori = errori + err.Key + ":" + err.Message + "\n"
		}
		return errors.Errorf(errori)
	}

	return nil
}
