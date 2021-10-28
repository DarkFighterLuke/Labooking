package utils

import (
	"Labooking/models"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

func RenderLayout(c *web.Controller) {
	c.Data["Ruolo"] = c.Ctx.Input.Session("ruolo")
	c.Data["NomeUtente"] = "Nome Cognome" // TODO: Implementare logica per mostrare il nome utente
	notifiche, err := GetNotificheByEmailPrivato(c)
	if err == nil {
		if len(notifiche) > 9 {
			c.Data["NumNotifiche"] = "9+"
		} else {
			c.Data["NumNotifiche"] = strconv.Itoa(len(notifiche))
		}
	}

	c.Layout = "dashboard/page_content_wrapper.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Navbar"] = "dashboard/navbar.tpl"
	c.LayoutSections["Sidebar"] = "dashboard/sidebar.tpl"
}

func GetNotificheByEmailPrivato(c *web.Controller) ([]*models.TestDiagnostico, error) {
	email := c.GetSession("email")
	if email == "" {
		return nil, fmt.Errorf("notifiche: %s", "email assente")
	}
	ruolo := c.GetSession("ruolo")
	if ruolo != "privato" {
		return nil, fmt.Errorf("notifiche: %s", "ruolo diverso da privato")
	}

	p := new(models.Privato)
	p.Email = email.(string)
	err := p.Seleziona("email")
	if err != nil {
		return nil, fmt.Errorf("notifiche: %s", err.Error())
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
