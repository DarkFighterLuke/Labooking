package models

import (
	"Labooking/models/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
)

func init() {
	orm.RegisterModel(new(Medico))
}

type Medico struct {
	IdMedico        int64   `orm:"pk;auto" form:"-"`
	Nome            string  `orm:"size(255)" form:"" valid:"Required" id:"nome-medico"`
	Cognome         string  `orm:"size(255)" form:"" valid:"Required" id:"cognome-medico"`
	CodiceFiscale   string  `orm:"size(16);unique" form:",,Codice Fiscale: " maxLength:"16" valid:"Required;Length(16)" id:"codice-fiscale-medico"`
	IndirizzoStudio string  `orm:"size(255)" form:",,Indirizzo studio: " valid:"Required" id:"indirizzo-medico"`
	Lat             float64 `orm:"digits(10);decimals(7)" form:"-"`
	Long            float64 `orm:"digits(10);decimals(7)" form:"-"`
	Prefisso        string  `orm:"size(6)" form:"-" valid:"Required"`
	Telefono        string  `orm:"size(10);unique" form:"" maxLength:"10" valid:"Required;Numeric;Length(10)" id:"telefono-medico"`
	Email           string  `orm:"size(255);unique" form:"" valid:"Required;Email" id:"email-medico"`
	Psw             string  `orm:"size(255)" form:"Password,password,Password: " valid:"Required" id:"password-medico"`
	ConfermaPsw     string  `orm:"-" form:"ConfermaPassword,password,Conferma password: " valid:"Required" id:"conferma-password-medico"`
	CodiceRegionale string  `orm:"size(255)" form:",,Codice regionale: " maxLength:"6" valid:"Required;Length(6)" id:"codice-regionale-medico"`
}

func (m *Medico) Aggiungi() (int64, error) {
	var err error
	m.Psw, err = utils.CryptSHA1(m.Psw)
	if err != nil {
		return -1, err
	}

	o := orm.NewOrm()
	newId, err := o.Insert(m)
	return newId, err
}

func (m *Medico) Seleziona(cols ...string) error {
	o := orm.NewOrm()
	err := o.Read(m, cols...)
	if err != nil {
		return err
	}
	return nil
}

func (m *Medico) Modifica() error {
	o := orm.NewOrm()
	_, err := o.Update(m)
	return err
}

func (m *Medico) Elimina(cols ...string) error {
	o := orm.NewOrm()
	err := m.Seleziona(cols...)
	_, err = o.Delete(m)
	return err
}

func (m *Medico) Valid(v *validation.Validation) {

	if m.Psw != m.ConfermaPsw {
		v.SetError("ConfermaPsw", "must be equal to Psw")
	}

	/*if !utils.IsPswValid(m.Psw) {
		v.SetError("Psw", "is not strong enough")
	}*/
}
