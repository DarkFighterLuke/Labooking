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
	rc.Data["FormOrganizzazione"] = &models.Organizzazione{}
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
		break
	case "medico":
		err := rc.registrazioneMedico()
		if err != nil {
			rc.Ctx.WriteString(err.Error())
			return
		}
		break
	case "laboratorio":
		err := rc.registrazioneLaboratorio()
		if err != nil {
			rc.Ctx.WriteString(err.Error())
			return
		}
		break
	case "organizzazione":
		err := rc.registrazioneOrganizzazione()
		if err != nil {
			rc.Ctx.WriteString(err.Error())
			return
		}
		break
	default:
		rc.Abort("400")
	}

	if rc.GetSession("ruolo") == "organizzazione" && rc.GetString("goToLogin") == "false" {
		rc.Redirect("/dashboard/dipendenti", http.StatusFound)
	} else if rc.GetSession("ruolo") == "medico" && rc.GetString("goToLogin") == "false" {
		rc.Redirect("/dashboard/pazienti", http.StatusFound)
	} else {
		rc.Redirect("/login", http.StatusFound)
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
	ll := models.LatLong{}
	err = models.RetrieveLatLong(p.Indirizzo, &ll)
	if err != nil {
		return err
	}
	p.Lat, err = strconv.ParseFloat(ll.Lat, 64)
	if err != nil {
		return err
	}
	p.Long, err = strconv.ParseFloat(ll.Long, 64)
	if err != nil {
		return err
	}
	if rc.GetSession("ruolo") == "organizzazione" {
		o := new(models.Organizzazione)
		o.Email = rc.GetSession("email").(string)
		err = o.Seleziona("email")
		if err != nil {
			return err
		}
		p.Organizzazione = o
		_, err = p.Aggiungi()
		if err != nil {
			return err
		}
	} else if rc.GetSession("ruolo") == "medico" {
		m := new(models.Medico)
		err = m.Seleziona("email")
		if err != nil {
			return err
		}
		p.Medico = m
		m.Email = rc.GetSession("email").(string)
		_, err = p.Aggiungi()
		if err != nil {
			return err
		}
	} else if p.Psw == "" {
		return fmt.Errorf("Psw Can not be empty")
	} else {
		_, err = p.Aggiungi()
		if err != nil {
			return err
		}
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
	ll := models.LatLong{}
	err = models.RetrieveLatLong(m.IndirizzoStudio, &ll)
	if err != nil {
		return err
	}
	m.Lat, err = strconv.ParseFloat(ll.Lat, 64)
	if err != nil {
		return err
	}
	m.Long, err = strconv.ParseFloat(ll.Long, 64)
	if err != nil {
		return err
	}
	_, err = m.Aggiungi()
	if err != nil {
		return err
	}
	return nil
}

func (rc *RegistrazioneController) registrazioneOrganizzazione() error {
	org := models.Organizzazione{}
	err := rc.ParseForm(&org)
	if err != nil {
		return err
	}
	org.Prefisso = rc.GetString("Prefisso")
	err = rc.validate(&org)
	if err != nil {
		return err

	}
	ll := models.LatLong{}
	err = models.RetrieveLatLong(org.Indirizzo, &ll)
	if err != nil {
		return err
	}
	org.Lat, err = strconv.ParseFloat(ll.Lat, 64)
	if err != nil {
		return err
	}
	org.Long, err = strconv.ParseFloat(ll.Long, 64)
	if err != nil {
		return err
	}
	_, err = org.Aggiungi()
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
	ll := models.LatLong{}
	err = models.RetrieveLatLong(l.Indirizzo, &ll)
	if err != nil {
		return err
	}
	l.Lat, err = strconv.ParseFloat(ll.Lat, 64)
	if err != nil {
		return err
	}
	l.Long, err = strconv.ParseFloat(ll.Long, 64)
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

	infoTest, err := rc.parseAndValidateInfoTest()
	if err != nil {
		return err
	}

	for _, it := range infoTest {
		it.IdLaboratorio = &models.Laboratorio{IdLaboratorio: id}
		_, err := it.Aggiungi()
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
	var orarioApertura = rc.GetString("orario-apertura-" + strconv.Itoa(i))
	var orarioChiusura = rc.GetString("orario-chiusura-" + strconv.Itoa(i))
	var giorno = rc.GetString("giorno-" + strconv.Itoa(i))
	for orarioApertura != "" || orarioChiusura != "" || giorno != "" {
		if orarioApertura == "" || orarioChiusura == "" || giorno == "" {
			break
		} else {
			oa, err := time.Parse("15:04", orarioApertura)
			if err != nil {
				return nil, err
			}
			oa = oa.AddDate(1, 0, 0)
			oc, err := time.Parse("15:04", orarioChiusura)
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
			orarioApertura = rc.GetString("orario-apertura-" + strconv.Itoa(i))
			orarioChiusura = rc.GetString("orario-chiusura-" + strconv.Itoa(i))
			giorno = rc.GetString("giorno-" + strconv.Itoa(i))
		}
	}
	return orari, nil
}

func (rc *RegistrazioneController) parseAndValidateInfoTest() ([]models.InfoTest, error) {
	var tipologiaTest = []string{"molecolare", "antigenico", "sierologico"}
	var ore = make([]int, 3)
	var minuti = make([]int, 3)
	var costo = make([]float64, 3)
	var effettua = make([]string, 3)

	var infoTest []models.InfoTest
	var err error

	for i := 0; i < 3; i++ {
		effettua[i] = rc.GetString(tipologiaTest[i] + "-effettua")
		if effettua[i] == "on" {
			ore[i], err = rc.GetInt(tipologiaTest[i] + "-ore")
			if err != nil {
				return nil, err
			}
			minuti[i], err = rc.GetInt(tipologiaTest[i] + "-minuti")
			if err != nil {
				return nil, err
			}
			costo[i], err = rc.GetFloat(tipologiaTest[i] + "-costo")
			if err != nil {
				return nil, err
			}
		}
	}

	for i := 0; i < 3; i++ {
		if ore[i] < 0 {
			return nil, fmt.Errorf("%s", "le ore necessarie ad analizzare un test diagnostico non possono essere inferiori a 0")
		}
		if minuti[i] != 0 && minuti[i] != 15 && minuti[i] != 30 && minuti[i] != 45 {
			return nil, fmt.Errorf("%s", "i minuti necessarie ad analizzare un test diagnostico non possono essere diversi dai valori prestabiliti")
		}
		if costo[i] < 0 || costo[i] > 9999.99 {
			return nil, fmt.Errorf("%s", "il costo di un test diagnostico non può essere negativo o superiore a 9999.99")
		}
		tempo, err := time.ParseDuration(strconv.Itoa(ore[i]) + "h" + strconv.Itoa(minuti[i]) + "m")
		if err != nil {
			return nil, err
		}

		if effettua[i] == "on" {
			it := models.InfoTest{
				TipologiaTest: tipologiaTest[i],
				Costo:         costo[i],
				Tempi:         int64(tempo.Seconds()),
			}
			infoTest = append(infoTest, it)
		}
	}

	return infoTest, nil
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
