package controllers

import "github.com/beego/beego/v2/server/web"

type HomepageController struct {
	web.Controller
}

func RenderDashboard(c *web.Controller) {

}

func (hc *HomepageController) Get() {
	hc.Data["UserType"] = hc.Ctx.Input.Session("ruolo")
	hc.Data["Title"] = "Attività recenti"
	hc.Data["NomeUtente"] = "Nome Cognome" // TODO: Implementare logica per mostrare il nome utente
	hc.Layout = "dashboard/page_content_wrapper.tpl"
	hc.LayoutSections = make(map[string]string)
	hc.LayoutSections["Navbar"] = "dashboard/navbar.tpl"
	hc.LayoutSections["Sidebar"] = "dashboard/sidebar.tpl"
	hc.TplName = "dashboard/home/home.tpl"
	hc.Data["NumNotifiche"] = "9+" // TODO: Implementare logica per il contatore notifiche
}
