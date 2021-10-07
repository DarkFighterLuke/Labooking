package routers

import (
	"Labooking/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.LandingPageController{})
	web.Router("/dashboard", &controllers.HomepageController{})

	//signp
	web.Router("/signup", &controllers.RegistrazioneController{})
	//web.Router("/recuperapassword", &controllers.RecuperoPasswordController{})
	//web.Router("/cambiapassword", &controllers.CambioPasswordController{})

	//login-logout
	web.Router("/login", &controllers.LoginController{})
	web.Router("/logout", &controllers.LogoutController{})
}
