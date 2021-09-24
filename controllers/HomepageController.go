package controllers

import "github.com/astaxie/beego"

type HomepageController struct {
	beego.Controller
}

func (c *HomepageController) Get() {
	switch c.GetString("page") {
	case "home":
		c.Data["Title"] = "Attività recenti"
		c.Data["TabName"] = "home"
		c.TplName = "dashboard/page_content_wrapper.tpl"
		break
	case "prenota":
		c.Data["Title"] = "Cerca laboratorio"
		c.Data["TabName"] = "prenota"
		c.TplName = "dashboard/page_content_wrapper.tpl"
		break
	case "referti":
		c.Data["Title"] = "I tuoi referti"
		c.Data["TabName"] = "referti"
		c.TplName = "dashboard/page_content_wrapper.tpl"
		break
	case "calendario":
		c.Data["Title"] = "Calendario"
		c.Data["TabName"] = "calendario"
		c.TplName = "dashboard/page_content_wrapper.tpl"
		break
	case "guida":
		c.Data["Title"] = "Guida ai test"
		c.Data["TabName"] = "guida"
		c.TplName = "dashboard/page_content_wrapper.tpl"
		break
	default:
		// TODO: Personalizzare pagina di errore con link alla dashboard
		c.Abort("404")
	}

}
