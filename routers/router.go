package routers

import (
	"Labooking/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	// pages
	web.Router("/", &controllers.LandingPageController{})
	web.Router("/dashboard/home", &controllers.HomepageController{})
	web.Router("/dashboard/prenota", &controllers.PrenotazioneController{})
	web.Router("/dashboard/guida", &controllers.GuidaController{})

	// signup
	web.Router("/signup", &controllers.RegistrazioneController{})
	web.Router("/recuperapassword", &controllers.RecuperoPasswordController{})
	web.Router("/cambiapassword", &controllers.CambioPasswordController{})

	// login-logout
	web.Router("/login", &controllers.LoginController{})
	web.Router("/logout", &controllers.LogoutController{})
}
