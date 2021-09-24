package routers

import (
	"Labooking/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/dashboard", &controllers.HomepageController{})
}
