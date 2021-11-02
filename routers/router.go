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
	web.Router("/dashboard/referti/download", &controllers.RefertoController{}, "get:DownloadReferto")
	web.Router("/dashboard/referti", &controllers.RefertoController{})
	web.Router("/dashboard/dipendenti", &controllers.PrivatoController{})
	web.Router("/dashboard/pazienti", &controllers.PrivatoController{})
	web.Router("/dashboard/dipendenti/aggiunta", &controllers.PrivatoController{}, "get:VisualizzaAggiunta")
	web.Router("/dashboard/pazienti/aggiunta", &controllers.PrivatoController{}, "get:VisualizzaAggiunta")

	web.Router("/api/ricerca", &api.RicercaLaboratorio{})
	web.Router("/api/notifiche", &api.Notifiche{}, "post:UpdateNotificheLette")

	// signup
	web.Router("/signup", &controllers.RegistrazioneController{})
	web.Router("/recuperapassword", &controllers.RecuperoPasswordController{})
	web.Router("/cambiapassword", &controllers.CambioPasswordController{})

	// login-logout
	web.Router("/login", &controllers.LoginController{})
	web.Router("/logout", &controllers.LogoutController{})
}
