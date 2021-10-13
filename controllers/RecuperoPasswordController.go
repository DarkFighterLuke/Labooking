package controllers

import (
	"Labooking/models"
	"Labooking/models/utils"
	"github.com/beego/beego/v2/server/web"
	"net/http"
)

type RecuperoPasswordController struct {
	web.Controller
}

func (rp *RecuperoPasswordController) Post() {
	ruolo := rp.GetString("idForm")

	switch ruolo {
	case "privato":
		p := new(models.Privato)
		p.Email = rp.GetString("email-privato")
		err := p.InviaLink()
		if err != nil {
			rp.Ctx.WriteString("invio mail: " + err.Error())
		}
		break
	case "laboratorio":
		l := new(models.Laboratorio)
		l.Email = rp.GetString("email-laboratorio")
		err := l.InviaLink()
		if err != nil {
			rp.Ctx.WriteString("errore nell'invio della mail di recupero")
		}
		break
	case "medico":
		m := new(models.Medico)
		m.Email = rp.GetString("email-medico")
		err := m.InviaLink()
		if err != nil {
			rp.Ctx.WriteString("errore nell'invio della mail di recupero")
		}
		break
	case "organizzazione":
		o := new(models.Organizzazione)
		o.Email = rp.GetString("email-organizzazione")
		err := o.InviaLink()
		if err != nil {
			rp.Ctx.WriteString("recupero password: " + err.Error())
		}
		break
	}
	rp.Redirect("/login", http.StatusFound)
}

func (rp *RecuperoPasswordController) Get() {
	rp.TplName = "login/recupero_password.tpl"
}

type CambioPasswordController struct {
	web.Controller
}

func (cp *CambioPasswordController) Get() {
	cp.TplName = "login/cambia_password.tpl"
}

func (cp *CambioPasswordController) Post() {
	r := new(models.RecuperoPassword)
	r.HashKey = cp.GetString("hash")
	err := r.CercaHash()
	if err != nil {
		cp.Ctx.WriteString("cambio password:" + err.Error())
	}

	newPsw := cp.GetString("password")
	confermaPsw := cp.GetString("conferma-password")
	if newPsw != confermaPsw {
		cp.Ctx.WriteString("cambio password: conferma password non andata a buon fine")
	}

	if r.Privato != nil {
		err := r.Privato.Seleziona("id_privato")
		if err != nil {
			cp.Ctx.WriteString("cambio password:" + err.Error())
		}

		newPsw, err = utils.CryptSHA1(newPsw)
		if err != nil {
			cp.Ctx.WriteString("cambio password:" + err.Error())
		}
		r.Privato.Psw = newPsw
		err = r.Privato.Modifica()
		if err != nil {
			cp.Ctx.WriteString("cambio password:" + err.Error())
		}
	} else if r.Medico != nil {
		err := r.Medico.Seleziona("id_medico")
		if err != nil {
			cp.Ctx.WriteString("cambio password:" + err.Error())
		}

		newPsw, err = utils.CryptSHA1(newPsw)
		if err != nil {
			cp.Ctx.WriteString("cambio password:" + err.Error())
		}
		r.Medico.Psw = newPsw
		err = r.Medico.Modifica()
		if err != nil {
			cp.Ctx.WriteString("cambio password:" + err.Error())
		}
	} else if r.Laboratorio != nil {
		err := r.Laboratorio.Seleziona("id_laboratorio")
		if err != nil {
			cp.Ctx.WriteString("cambio password:" + err.Error())
		}

		newPsw, err = utils.CryptSHA1(newPsw)
		if err != nil {
			cp.Ctx.WriteString("cambio password:" + err.Error())
		}
		r.Laboratorio.Psw = newPsw
		err = r.Laboratorio.Modifica()
		if err != nil {
			cp.Ctx.WriteString("cambio password:" + err.Error())
		}
	} else if r.Organizzazione != nil {
		err := r.Organizzazione.Seleziona("id_organizzazione")
		if err != nil {
			cp.Ctx.WriteString("cambio password:" + err.Error())
		}

		newPsw, err = utils.CryptSHA1(newPsw)
		if err != nil {
			cp.Ctx.WriteString("cambio password:" + err.Error())
		}
		r.Organizzazione.Psw = newPsw
		err = r.Organizzazione.Modifica()
		if err != nil {
			cp.Ctx.WriteString("cambio password:" + err.Error())
		}
	}

	err = r.Elimina()
	if err != nil {
		cp.Ctx.WriteString("cambio password:" + err.Error())
	}

	cp.TplName = "login/login.tpl"
}
