package routers

import (
	"Labooking/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.LandingPageController{})
	web.Router("/dashboard", &controllers.HomepageController{})
	web.Router("/signup", &controllers.RegistrazioneController{})
}
