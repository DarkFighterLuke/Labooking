package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(Privato))
}

type Privato struct {
	IdPrivato              int       `orm:"pk;auto" form:"-"`
	Nome                   string    `orm:"size(255)" form:""`
	Cognome                string    `orm:"size(255)" form:""`
	CodiceFiscale          string    `orm:"size(16);unique" form:",,Codice Fiscale: " maxLength:"16"`
	NumeroTesseraSanitaria string    `orm:"size(20);unique" form:",,Numero Tessera Sanitaria: " maxLength:"20"`
	Citta                  string    `orm:"size(255)" form:""`
	Cap                    string    `orm:"size(5)" form:",,CAP: " maxLength:"5"`
	Via                    string    `orm:"size(255)" form:",,Via/Piazza: "`
	Civico                 int       `orm:"digits(4)" form:"" maxLength:"4"`
	Prefisso               string    `orm:"size(2)" form:"" maxLength:"2"`
	Telefono               string    `orm:"size(10);unique" form:"" maxLength:"10"`
	Email                  string    `orm:"size(255);unique" form:""`
	EmailConfermata        bool      `form:"-"`
	Psw                    string    `orm:"size(255)" form:"Password,password,Password: "`
	DataNascita            time.Time `orm:"type(date)"`
	Medico                 *Medico   `orm:"rel(fk)" form:"-"`
}

func (p *Privato) Aggiungi() error {
	o := orm.NewOrm()
	_, err := o.Insert(p)
	return err
}

func (p *Privato) Seleziona() (interface{}, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(p)
	privati := make([]Privato, 0)
	_, err := qs.All(privati)
	if err != nil {
		return nil, err
	}
	return privati, nil
}

func (p *Privato) Modifica() error {
	o := orm.NewOrm()
	_, err := o.Update(p)
	return err
}

func (p *Privato) Elimina() error {
	o := orm.NewOrm()
	_, err := o.Delete(p)
	return err
}
