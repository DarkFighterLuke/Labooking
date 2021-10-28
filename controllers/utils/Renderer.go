package utils

import (
	"Labooking/models"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

func RenderLayout(c *web.Controller) {
	c.Data["Ruolo"] = c.Ctx.Input.Session("ruolo")
	nomeUtente, err := GetNomeUtente(c)
	c.Data["NomeUtente"] = nomeUtente
	notifiche, err := GetNotificheByEmailPrivato(c)
	if err == nil {
		if len(notifiche) > 9 {
			c.Data["NumNotifiche"] = "9+"
		} else {
			c.Data["NumNotifiche"] = strconv.Itoa(len(notifiche))
		}
		c.Data["Notifiche"] = notifiche
	}

	c.Layout = "dashboard/page_content_wrapper.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Navbar"] = "dashboard/navbar.tpl"
	c.LayoutSections["Sidebar"] = "dashboard/sidebar.tpl"
}

func GetNotificheByEmailPrivato(c *web.Controller) ([]*models.TestDiagnostico, error) {
	email := c.GetSession("email")
	ruolo := c.GetSession("ruolo")
	if ruolo != "privato" {
		return nil, fmt.Errorf("getnotifiche: %s", "ruolo diverso da privato")
	}

	p := new(models.Privato)
	p.Email = email.(string)
	err := p.Seleziona("email")
	if err != nil {
		return nil, fmt.Errorf("getnotifiche: %s", err.Error())
	}

	td := new(models.TestDiagnostico)
	td.Privato = p
	td.Stato = "notificato"
	testDiagnostici, err := td.SelezionaAllTestsByPrivatoStato()
	if err != nil {
		return nil, fmt.Errorf("notifiche: %s", err.Error())
	}

	return testDiagnostici, err
}

func GetNomeUtente(c *web.Controller) (string, error) {
	email := c.GetSession("email")
	ruolo := c.GetSession("ruolo")

	if ruolo == "privato" {
		p := new(models.Privato)
		p.Email = email.(string)
		err := p.Seleziona("email")
		if err != nil {
			return "", fmt.Errorf("getnomeutente: %s", err.Error())
		}
		return p.Nome + " " + p.Cognome, err
	} else if ruolo == "medico" {
		m := new(models.Medico)
		m.Email = email.(string)
		err := m.Seleziona("email")
		if err != nil {
			return "", fmt.Errorf("getnomeutente: %s", err.Error())
		}
		return m.Nome + " " + m.Cognome, err
	} else if ruolo == "organizzazione" {
		o := new(models.Organizzazione)
		o.Email = email.(string)
		err := o.Seleziona("email")
		if err != nil {
			return "", fmt.Errorf("getnomeutente: %s", err.Error())
		}
		return o.RagioneSociale, err
	} else if ruolo == "laboratorio" {
		l := new(models.Laboratorio)
		l.Email = email.(string)
		err := l.Seleziona("email")
		if err != nil {
			return "", fmt.Errorf("getnomeutente: %s", err.Error())
		}
		return l.Nome, err
	} else {
		return "", fmt.Errorf("getnomeutente: %s", "ruolo sconosciuto")
	}
}
