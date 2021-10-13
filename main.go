package main

import (
	"Labooking/models/utils"
	_ "Labooking/routers"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/session"
	_ "github.com/go-sql-driver/mysql"
	"log"
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

	//timer
	go utils.Timer()

	//timer
	//go utils.Timer()

	//database config
	driver, err := web.AppConfig.String("mysqldriver")
	if err != nil {
		log.Fatalln(err)
	}

	err = orm.RegisterDriver(driver, orm.DRMySQL)
	if err != nil {
		log.Fatalln(err.Error())
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
		"?charset=utf8&loc=Europe%2FRome"
	err = orm.RegisterDataBase("default", driver, connectionString)
	if err != nil {
		log.Fatalln(err.Error())
	}

	//Configure session
	sessionconf := &session.ManagerConfig{
		CookieName:              "LabookingSession",
		SessionNameInHTTPHeader: "LabookingSession",
		Gclifetime:              3600,
		Maxlifetime:             3600,
	}
	web.GlobalSessions, err = session.NewManager("memory", sessionconf)
	if err != nil {
		log.Fatalln(err)
	}
	go web.GlobalSessions.GC()

	//filters
	web.InsertFilter("/dashboard/*", web.BeforeRouter, filterUser)

}

//filter user configuration
var filterUser = func(ctx *context.Context) {
	email := ctx.Input.Session("email")
	ruolo := ctx.Input.Session("ruolo")

	if email == nil || ruolo == nil {
		ctx.Redirect(302, "/login")
	}
}
