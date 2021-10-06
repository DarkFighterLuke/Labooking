package models

import (
	"Labooking/models/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	"time"
)

func init() {
	orm.RegisterModel(new(Privato))
}

type Privato struct {
	IdPrivato              int64     `orm:"pk;auto" form:"-"`
	Nome                   string    `orm:"size(255)" form:"" valid:"Required" id:"nome-privato"`
	Cognome                string    `orm:"size(255)" form:"" valid:"Required" id:"cognome-privato"`
	CodiceFiscale          string    `orm:"size(16);unique" form:",,Codice Fiscale: " maxLength:"16" valid:"Required;Length(16)" id:"codice-fiscale-privato"`
	NumeroTesseraSanitaria string    `orm:"size(20);unique" form:",,Numero Tessera Sanitaria: " maxLength:"20" valid:"Required;Length(20)" id:"numero-tessera-sanitaria-privato"`
	Citta                  string    `orm:"size(255)" form:"" valid:"Required" id:"citta-privato"`
	Cap                    string    `orm:"size(5)" form:",,CAP: " maxLength:"5" valid:"Required;Length(5)" id:"cap-privato"`
	Via                    string    `orm:"size(255)" form:",,Via/Piazza: " valid:"Required" id:"via-privato"`
	Civico                 int       `orm:"digits(4)" form:"" maxLength:"4" valid:"Required;Range(1, 9999)" id:"civico-privato"`
	Prefisso               string    `orm:"size(6)" form:"-" valid:"Required"`
	Telefono               string    `orm:"size(10);unique" form:"" maxLength:"10" valid:"Required;Numeric;Length(10)" id:"telefono-privato"`
	Email                  string    `orm:"size(255);unique" form:"" valid:"Required;Email" id:"email-privato"`
	EmailConfermata        bool      `form:"-"`
	Psw                    string    `orm:"size(255)" form:"Password,password,Password: " valid:"Required" id:"password-privato"`
	ConfermaPsw            string    `orm:"-" form:"ConfermaPassword,password,Conferma password: " valid:"Required" id:"conferma-password-privato"`
	DataNascita            time.Time `orm:"type(date)" valid:"Required"`
	Medico                 *Medico   `orm:"rel(fk);null;on_delete(set_null);column(medico)" form:"-"`
}

func (p *Privato) Aggiungi() (int64, error) {
	var found = false
	ptemp := *p
	err := ptemp.Seleziona("numero_tessera_sanitaria")
	if err == nil {
		found = true
		p.IdPrivato = ptemp.IdPrivato
		p.Nome = ptemp.Nome
		p.Cognome = ptemp.Cognome
		p.CodiceFiscale = ptemp.CodiceFiscale
		p.NumeroTesseraSanitaria = ptemp.NumeroTesseraSanitaria
		p.DataNascita = ptemp.DataNascita
		p.Medico = ptemp.Medico
	}

	p.Psw, err = utils.CryptSHA1(p.Psw)
	if err != nil {
		return -1, err
	}

	o := orm.NewOrm()
	if found {
		_, err = o.Update(p)
		return int64(p.IdPrivato), err
	} else {
		newId, err := o.Insert(p)
		return newId, err
	}
}

func (p *Privato) Seleziona(cols ...string) error {
	o := orm.NewOrm()
	err := o.Read(p, cols...)
	if err != nil {
		return err
	}
	return nil
}

func (p *Privato) Modifica() error {
	o := orm.NewOrm()
	_, err := o.Update(p)
	return err
}

func (p *Privato) Elimina(cols ...string) error {
	o := orm.NewOrm()
	err := p.Seleziona(cols...)
	_, err = o.Delete(p)
	return err
}

func (p *Privato) Valid(v *validation.Validation) {
	//data
	if !utils.IsAdult(p.DataNascita) {
		v.SetError("DataNascita", "must be eighteen years old")
	}

	if p.Psw != p.ConfermaPsw {
		v.SetError("ConfermaPsw", "must be equal to Psw")
	}

	/*if !utils.IsPswValid(p.Psw) {
		v.SetError("Psw", "is not strong enough")
	}*/
}
