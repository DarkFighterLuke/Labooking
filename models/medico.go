package models

import (
	"Labooking/models/utils"
	"encoding/hex"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
)

func init() {
	orm.RegisterModel(new(Medico))
}

type Medico struct {
	IdMedico        int    `orm:"pk;auto" form:"-"`
	Nome            string `orm:"size(255)" form:"" valid:"Required"`
	Cognome         string `orm:"size(255)" form:"" valid:"Required"`
	CodiceFiscale   string `orm:"size(16);unique" form:",,Codice Fiscale: " maxLength:"16" valid:"Required;Length(16)"`
	CittaStudio     string `orm:"size(255)" form:",,Città dello studio: " valid:"Required"`
	CapStudio       string `orm:"size(5)" form:",,CAP dello studio: " maxLength:"5" valid:"Required;Length(5)"`
	ViaStudio       string `orm:"size(255)" form:",,Via/Piazza dello studio: " valid:"Required"`
	CivicoStudio    int    `orm:"digits(4)" form:",,Civico dello studio: " maxLength:"4" valid:"Required;Range(1, 9999)"`
	Prefisso        string `orm:"size(2)" form:"-" maxLength:"2" valid:"Required;Length(2)"`
	Telefono        string `orm:"size(10);unique" form:"" maxLength:"10" valid:"Required;Numeric;Length(10)"`
	Email           string `orm:"size(255);unique" form:"" valid:"Required;Email"`
	EmailConfermata bool   `form:"-"`
	Psw             string `orm:"size(255)" form:"Password,password,Password: " valid:"Required"`
	ConfermaPsw     string `orm:"-" form:"ConfermaPassword,password,Conferma password: " valid:"Required"`
	CodiceRegionale string `orm:"size(255)" form:",,Codice regionale: " maxLength:"7" valid:"Required;Length(7)"`
}

func (m *Medico) Aggiungi() error {
	bytePsw, err := utils.CryptSHA1(m.Psw)
	if err != nil {
		return err
	}
	m.Psw = hex.EncodeToString(bytePsw)

	o := orm.NewOrm()
	_, err = o.Insert(m)
	return err
}

func (m *Medico) Seleziona() (interface{}, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(m)
	medici := make([]Medico, 0)
	_, err := qs.All(medici)
	if err != nil {
		return nil, err
	}
	return medici, nil
}

func (m *Medico) Modifica() error {
	o := orm.NewOrm()
	_, err := o.Update(m)
	return err
}

func (m *Medico) Elimina() error {
	o := orm.NewOrm()
	_, err := o.Delete(m)
	return err
}

func (m *Medico) Valid(v *validation.Validation) {

	if m.Psw != m.ConfermaPsw {
		v.SetError("ConfermaPsw", "must be equal to Psw")
	}

	if !utils.IsPswValid(m.Psw) {
		v.SetError("Psw", "is not strong enough")
	}
}
