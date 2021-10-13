package models

import (
	"Labooking/models/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
)

func init() {
	orm.RegisterModel(new(Laboratorio))
}

type Laboratorio struct {
	IdLaboratorio int64            `orm:"pk;auto" form:"-"`
	Nome          string           `orm:"size(255)" form:"" valid:"Required" id:"nome-laboratorio"`
	PartitaIva    string           `orm:"size(11);unique" form:",,Partita Iva: " maxLength:"11" valid:"Required;Length(11)" id:"partita-iva-laboratorio"`
	Indirizzo     string           `orm:"size(255)" form:"" valid:"Required" id:"indirizzo-laboratorio"`
	Lat           float64          `orm:"digits(10);decimals(7)" form:"-"`
	Long          float64          `orm:"digits(10);decimals(7)" form:"-"`
	Prefisso      string           `orm:"size(6)" form:"-" valid:"Required"`
	Telefono      string           `orm:"size(10);unique" form:"" maxLength:"10" valid:"Required;Numeric;Length(10)" id:"telefono-laboratorio"`
	Email         string           `orm:"size(255);unique" form:"" valid:"Required;Email" id:"email-laboratorio"`
	Psw           string           `orm:"size(255)" form:"Password,password,Password: " valid:"Required" id:"password-laboratorio"`
	ConfermaPsw   string           `orm:"-" form:"ConfermaPassword,password,Conferma password: " valid:"Required" id:"conferma-password-laboratorio"`
	Test          []*InfoTest      `orm:"reverse(many)" form:"-"`
	Orari         []*OrariApertura `orm:"reverse(many)" form:"-"`
}

func (l *Laboratorio) Aggiungi() (int64, error) {
	var err error
	l.Psw, err = utils.CryptSHA1(l.Psw)
	if err != nil {
		return -1, err
	}

	o := orm.NewOrm()
	newId, err := o.Insert(l)
	return newId, err
}

func (l *Laboratorio) Seleziona(cols ...string) error {
	o := orm.NewOrm()
	err := o.Read(l, cols...)
	if err != nil {
		return err
	}
	return nil
}

func (l *Laboratorio) Modifica() error {
	o := orm.NewOrm()
	_, err := o.Update(l)
	return err
}

func (l *Laboratorio) Elimina(cols ...string) error {
	o := orm.NewOrm()
	err := l.Seleziona(cols...)
	_, err = o.Delete(l)
	return err
}

func (l *Laboratorio) Valid(v *validation.Validation) {

	if l.Psw != l.ConfermaPsw {
		v.SetError("ConfermaPsw", "must be equal to Psw")
	}

	/*if !utils.IsPswValid(l.Psw) {
		v.SetError("Psw", "is not strong enough")
	}*/
}
