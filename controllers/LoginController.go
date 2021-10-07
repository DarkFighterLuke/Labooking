package controllers

import (
	"Labooking/models"
	"Labooking/models/utils"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"net/http"
)

type LoginController struct {
	web.Controller
}

func (li *LoginController) Get() {
	email := fmt.Sprint(li.GetSession("email"))
	ruolo := fmt.Sprint(li.GetSession("ruolo"))

	//Controlla se il login è stato gia effettuato
	ok, _ := checkUserExistence(email, ruolo)

	if ok {
		li.Redirect("/dashboard?page=home", http.StatusPermanentRedirect)
	} else {
		li.TplName = "login/login.tpl"
	}

}

func (li *LoginController) Post() {
	ruolo := li.GetString("idForm")
	var password string
	var email string
	switch ruolo {
	case "privato":
		email = li.GetString("email-privato")
		password = li.GetString("password-privato")
		break
	case "laboratorio":
		email = li.GetString("email-laboratorio")
		password = li.GetString("password-laboratorio")
		break
	case "medico":
		email = li.GetString("email-medico")
		password = li.GetString("password-medico")
		break
	}

	li.StartSession()

	if email == "" || ruolo == "" || password == "" {
		li.TplName = "login/login.tpl"
		li.Data["errmsg"] = "login failed: inserire nome utente o password"
		return
	}

	ok, psw := checkUserExistence(email, ruolo)
	if !ok {
		li.TplName = "login/login.tpl"
		li.Data["errmsg"] = "login failed: utente non trovato"
		return
	}

	passwordSha1, err := utils.CryptSHA1(password)
	if err != nil {
		li.Abort("500")
		return
	}

	if passwordSha1 != psw {
		li.TplName = "login/login.tpl"
		li.Data["errmsg"] = "login failed: wrong password"
		return
	}

	err = li.SetSession("email", email)
	err = li.SetSession("ruolo", ruolo)
	if err != nil {
		li.TplName = "login/login.tpl"
		li.Data["errmsg"] = err.Error()
		return
	}
	li.Redirect("/dashboard?page=home", http.StatusFound)
}

func checkUserExistence(email, ruolo string) (bool, string) {
	switch ruolo {
	case "privato":
		p := new(models.Privato)
		p.Email = fmt.Sprint(email)

		err := p.Seleziona("email")
		if err != nil {
			return false, ""
		}
		return true, p.Psw
		break
	case "laboratorio":
		l := new(models.Laboratorio)
		l.Email = fmt.Sprint(email)

		err := l.Seleziona("email")
		if err != nil {
			return false, ""
		}
		return true, l.Psw
		break
	case "medico":
		m := new(models.Medico)
		m.Email = fmt.Sprint(email)

		err := m.Seleziona("email")
		if err != nil {
			return false, ""
		}
		return true, m.Psw
		break
	}
	return false, ""
}
