package main

import (
	_ "Labooking/routers"
	"github.com/astaxie/beego"
)

func main() {
	web.Run()
}

func init() {
	//static files
	web.SetStaticPath("/css", "static/css")
	web.SetStaticPath("/fonts", "static/fonts")
	web.SetStaticPath("/img", "static/img")
	web.SetStaticPath("/js", "static/js")

	//database config
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		log.Fatalln(err.Error())
	}

	driver, err := web.AppConfig.String("mysqldriver")
	if err != nil {
		log.Fatalln(err)
	}
	mysqluser, err := web.AppConfig.String("mysqluser")
	if err != nil {
		log.Fatalln(err)
	}
	mysqlpsw, err := web.AppConfig.String("mysqlpsw")
	if err != nil {
		log.Fatalln(err)
	}
	mysqlurl, err := web.AppConfig.String("mysqlurl")
	if err != nil {
		log.Fatalln(err)
	}
	mysqlport, err := web.AppConfig.String("mysqlport")
	if err != nil {
		log.Fatalln(err)
	}
	mysqldbname, err := web.AppConfig.String("mysqldbname")
	if err != nil {
		log.Fatalln(err)
	}
	connectionString := mysqluser + ":" + mysqlpsw + "@tcp(" +
		mysqlurl + ":" + mysqlport + ")/" + mysqldbname +
		"?charset=utf8"
	err = orm.RegisterDataBase("default", driver, connectionString)
	if err != nil {
		log.Fatalln(err.Error())
	}

}
