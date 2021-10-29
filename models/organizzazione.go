package models

import (
	"Labooking/models/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
)

func init() {
	orm.RegisterModel(new(Organizzazione))
}

type Organizzazione struct {
	IdOrganizzazione int        `orm:"pk;auto" form:"-"`
	RagioneSociale   string     `orm:"size(255)" form:"" valid:"Required" id:"ragione-sociale-organizzazione"`
	PartitaIva       string     `orm:"size(11);unique" form:",,Partita Iva: " maxLength:"11" valid:"Required;Length(11)" id:"partita-iva-organizzazione"`
	Indirizzo        string     `orm:"size(255)" form:"" valid:"Required" id:"indirizzo-organizzazione"`
	Lat              float64    `orm:"digits(10);decimals(7)" form:"-"`
	Long             float64    `orm:"digits(10);decimals(7)" form:"-"`
	Prefisso         string     `orm:"size(6)" form:"-" valid:"Required"`
	Telefono         string     `orm:"size(10);unique" form:"" maxLength:"10" valid:"Required;Numeric;Length(10)" id:"telefono-organizzazione"`
	Email            string     `orm:"size(255);unique" form:"" valid:"Required;Email" id:"email-organizzazione"`
	Psw              string     `orm:"size(255)" form:"Password,password,Password: " valid:"Required" id:"password-organizzazione"`
	ConfermaPsw      string     `orm:"-" form:"ConfermaPassword,password,Conferma password: " valid:"Required" id:"conferma-password-organizzazione"`
	Dipendenti       []*Privato `orm:"reverse(many)"`
}

func (org *Organizzazione) Aggiungi() (int64, error) {
	var err error
	org.Psw, err = utils.CryptSHA1(org.Psw)
	if err != nil {
		return -1, err
	}

	o := orm.NewOrm()
	newId, err := o.Insert(org)
	if err != nil {
		return -1, err
	}
	return newId, nil
}

func (org *Organizzazione) Seleziona(cols ...string) error {
	o := orm.NewOrm()
	err := o.Read(org, cols...)
	if err != nil {
		return err
	}
	return nil
}

func (org *Organizzazione) Modifica() error {
	o := orm.NewOrm()
	_, err := o.Update(org)
	return err
}

func (org *Organizzazione) Elimina(cols ...string) error {
	o := orm.NewOrm()
	err := org.Seleziona(cols...)
	_, err = o.Delete(org)
	return err
}

func (org *Organizzazione) GetDipendenti() ([]*Privato, error) {
	o := orm.NewOrm()
	_, err := o.LoadRelated(org, "Dipendenti")
	if err != nil {
		return nil, err
	}

	return org.Dipendenti, err
}

func (org *Organizzazione) Valid(v *validation.Validation) {
	if org.Psw != org.ConfermaPsw {
		v.SetError("ConfermaPsw", "must be equal to Psw")
	}
}
