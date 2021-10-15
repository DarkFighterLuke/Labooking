package test

import (
	"Labooking/models"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
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
	err = orm.RegisterDataBase("default", driver, connectionString)
	if err != nil {
		log.Fatalln(err.Error())
	}

}

func TestAggiungi(t *testing.T) {
	p := new(models.Privato)

	p.Nome = "Tizio"
	p.Cognome = "Caio"
	p.CodiceFiscale = "123456789qwertyu"
	p.NumeroTesseraSanitaria = "12345678912345678911"
	p.Citta = "Roma"
	p.Cap = "00000"
	p.Via = "RRRRR"
	p.Civico = 123
	p.Prefisso = "39"
	p.Telefono = "234234234"
	p.Email = "tizio.caio@vvv.it"
	p.EmailConfermata = false
	p.Psw = "qwerty"
	p.DataNascita = time.Date(2000, 4, 3, 12, 30, 0, 0, time.UTC)
	p.Medico = nil

	err := p.Aggiungi()

	if err != nil {
		t.Errorf("err is %v; want nil", err)
	}
}

func TestSeleziona(t *testing.T) {
	p := new(models.Privato)

	p.Nome = "Tizio"
	p.Cognome = "Caio"

	err := p.Seleziona("nome", "cognome")
	if err != nil {
		t.Errorf("err is %v, want nil", err)
	}
	fmt.Printf("%#v", p)
	fmt.Printf("%#v", p.Medico)
}

func TestModifica(t *testing.T) {
	p := new(models.Privato)

	p.Nome = "Tizio"

	err := p.Seleziona("nome")
	if err != nil {
		t.Errorf("err is %v, want nil", err)
	}
	p.Medico = new(models.Medico)
	p.Medico.IdMedico = 1

	err = p.Modifica()
	if err != nil {
		t.Errorf("err is %v, want nil", err)
	}
}

func TestElimina(t *testing.T) {
	p := new(models.Privato)

	p.Nome = "Tizio"

	err := p.Elimina("nome")
	if err != nil {
		t.Errorf("err is %v, want nil", err)
	}
}
