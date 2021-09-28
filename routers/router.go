package routers

import (
	"Labooking/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.LandingPageController{})
	beego.Router("/dashboard", &controllers.HomepageController{})
}
