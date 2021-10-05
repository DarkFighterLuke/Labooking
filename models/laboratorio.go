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
	IdLaboratorio   int              `orm:"pk;auto" form:"-"`
	Nome            string           `orm:"size(255)" form:"" valid:"Required" id:"nome-laboratorio"`
	PartitaIva      string           `orm:"size(11);unique" form:",,Partita Iva: " maxLength:"11" valid:"Required;Length(11)" id:"partita-iva-laboratorio"`
	Citta           string           `orm:"size(255)" form:"" valid:"Required" id:"citta-laboratorio"`
	Cap             string           `orm:"size(5)" form:",,CAP: " maxLength:"5" valid:"Required;Length(5)" id:"cap-laboratorio"`
	Via             string           `orm:"size(255)" form:",,Via/Piazza: " valid:"Required" id:"via-laboratorio"`
	Civico          int              `orm:"digits(4)" form:"" maxLength:"4" valid:"Required;Range(1,9999)" id:"civico-laboratorio"`
	Prefisso        string           `orm:"size(6)" form:"-" valid:"Required"`
	Telefono        string           `orm:"size(10);unique" form:"" maxLength:"10" valid:"Required;Numeric;Length(10)" id:"telefono-laboratorio"`
	Email           string           `orm:"size(255);unique" form:"" valid:"Required;Email" id:"email-laboratorio"`
	EmailConfermata bool             `form:"-"`
	Psw             string           `orm:"size(255)" form:"Password,password,Password: " valid:"Required" id:"password-laboratorio"`
	ConfermaPsw     string           `orm:"-" form:"ConfermaPassword,password,Conferma password: " valid:"Required" id:"conferma-password-laboratorio"`
	Test            []*InfoTest      `orm:"reverse(many)" form:"-"`
	Orari           []*OrariApertura `orm:"reverse(many)" form:"-"`
}

func (l *Laboratorio) Aggiungi() error {
	var err error
	l.Psw, err = utils.CryptSHA1(l.Psw)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	_, err = o.Insert(l)
	return err
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
