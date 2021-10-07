package controllers

import "github.com/beego/beego/v2/server/web"

type HomepageController struct {
	web.Controller
}

func (c *HomepageController) Get() {
	switch c.GetString("page") {
	case "home":
		c.Data["Title"] = "Attività recenti"
		c.Data["NomeUtente"] = "Nome Cognome" // TODO: Implementare logica per mostrare il nome utente
		c.Data["TabName"] = "home"
		c.Data["NumNotifiche"] = "9+" // TODO: Implementare logica per il contatore notifiche
		break
	case "prenota":
		c.Data["Title"] = "Cerca laboratorio"
		c.Data["NomeUtente"] = "Nome Cognome" // TODO: Implementare logica per mostrare il nome utente
		c.Data["TabName"] = "prenota"
		c.Data["NumNotifiche"] = "9+" // TODO: Implementare logica per il contatore notifiche
		break
	case "referti":
		c.Data["Title"] = "I tuoi referti"
		c.Data["NomeUtente"] = "Nome Cognome" // TODO: Implementare logica per mostrare il nome utente
		c.Data["TabName"] = "referti"
		c.Data["NumNotifiche"] = "9+" // TODO: Implementare logica per il contatore notifiche
		break
	case "calendario":
		c.Data["Title"] = "Calendario"
		c.Data["NomeUtente"] = "Nome Cognome" // TODO: Implementare logica per mostrare il nome utente
		c.Data["TabName"] = "calendario"
		c.Data["NumNotifiche"] = "9+" // TODO: Implementare logica per il contatore notifiche
		break
	case "guida":
		c.Data["Title"] = "Guida ai test"
		c.Data["NomeUtente"] = "Nome Cognome" // TODO: Implementare logica per mostrare il nome utente
		c.Data["TabName"] = "guida"
		c.Data["NumNotifiche"] = "9+" // TODO: Implementare logica per il contatore notifiche
		break
	default:
		// TODO: Personalizzare pagina di errore con link alla dashboard
		c.Abort("404")
		return
	}
	c.TplName = "dashboard/page_content_wrapper.tpl"
}
