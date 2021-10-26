package routers

import (
	"Labooking/controllers"
	"Labooking/controllers/api"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	// pages
	web.Router("/", &controllers.LandingPageController{})
	web.Router("/dashboard/home", &controllers.HomepageController{})
	web.Router("/dashboard/prenota", &controllers.PrenotazioneController{})
	web.Router("/dashboard/guida", &controllers.GuidaController{})
	web.Router("/dashboard/laboratorio", &controllers.LaboratorioController{})
	web.Router("/dashboard/prenotazioni", &controllers.PrenotazioneController{}, "get:VisualizzaPrenotazioni")

	web.Router("/api/ricerca", &api.RicercaLaboratorio{})

	// signup
	web.Router("/signup", &controllers.RegistrazioneController{})
	web.Router("/recuperapassword", &controllers.RecuperoPasswordController{})
	web.Router("/cambiapassword", &controllers.CambioPasswordController{})

	// login-logout
	web.Router("/login", &controllers.LoginController{})
	web.Router("/logout", &controllers.LogoutController{})
}
