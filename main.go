package main

import (
	_ "Labooking/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

func init() {
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/fonts", "static/fonts")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/js", "static/js")
}
