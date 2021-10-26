package test

import (
	"Labooking/models"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"log"
	"testing"
	"time"
)

func init() {
	//static files
	web.SetStaticPath("/css", "static/css")
	web.SetStaticPath("/fonts", "static/fonts")
	web.SetStaticPath("/img", "static/img")
	web.SetStaticPath("/js", "static/js")

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
		"?charset=utf8"
	err = orm.RegisterDataBase("test", driver, connectionString)
	if err != nil {
		log.Fatalln(err.Error())
	}

}

func TestFiltraLaboratorio(t *testing.T) {
	var labs []models.Laboratorio
	tipi := []string{"molecolare", "antigenico"}
	ora_inizio := time.Date(0, 0, 0, 10, 0, 0, 0, time.UTC)
	ora_fine := time.Date(0, 0, 0, 11, 0, 0, 0, time.UTC)
	err := models.FiltraLaboratori(&labs, 173701, tipi, 70, ora_inizio, ora_fine, "lunedi")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(labs)
}
