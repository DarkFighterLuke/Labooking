package controllers

import "github.com/beego/beego/v2/server/web"

type GuidaController struct {
	web.Controller
}

func (gc *GuidaController) Get() {
	gc.Data["UserType"] = gc.Ctx.Input.Session("ruolo")
	gc.Data["Title"] = "Guida ai test"
	gc.Data["NomeUtente"] = "Nome Cognome" // TODO: Implementare logica per mostrare il nome utente
	gc.Layout = "dashboard/page_content_wrapper.tpl"
	gc.LayoutSections = make(map[string]string)
	gc.LayoutSections["Navbar"] = "dashboard/navbar.tpl"
	gc.LayoutSections["Sidebar"] = "dashboard/sidebar.tpl"
	gc.TplName = "dashboard/guida/guida.tpl"
	gc.Data["NumNotifiche"] = "9+" // TODO: Implementare logica per il contatore notifiche
}
