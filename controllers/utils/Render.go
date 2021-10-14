package utils

import "github.com/beego/beego/v2/server/web"

func RenderLayout(c *web.Controller) {
	c.Data["UserType"] = c.Ctx.Input.Session("ruolo")
	c.Data["NomeUtente"] = "Nome Cognome" // TODO: Implementare logica per mostrare il nome utente
	c.Data["NumNotifiche"] = "9+"         // TODO: Implementare logica per il contatore notifiche
	c.Layout = "dashboard/page_content_wrapper.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Navbar"] = "dashboard/navbar.tpl"
	c.LayoutSections["Sidebar"] = "dashboard/sidebar.tpl"
}
