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
	Nome                   string    `orm:"size(255)" form:"" valid:"Required;MinSize(1);MaxSize(255);Alpha" id:"nome-privato"`
	Cognome                string    `orm:"size(255)" form:"" valid:"Required;MinSize(1);MaxSize(255);Alpha" id:"cognome-privato"`
	CodiceFiscale          string    `orm:"size(16);unique" form:",,Codice Fiscale: " maxLength:"16" valid:"Required;Length(16)" id:"codice-fiscale-privato"`
	NumeroTesseraSanitaria string    `orm:"size(20);unique" form:",,Numero Tessera Sanitaria: " maxLength:"20" valid:"Required;Length(20)" id:"numero-tessera-sanitaria-privato"`
	Citta                  string    `orm:"size(255)" form:"" valid:"Required;MaxSize(255);Alpha" id:"citta-privato"`
	Cap                    string    `orm:"size(5)" form:",,CAP: " maxLength:"5" valid:"Required;Length(5);Numeric" id:"cap-privato"`
	Via                    string    `orm:"size(255)" form:",,Via/Piazza: " valid:"Required;MinSize(1);MaxSize(255);AlphaNumeric" id:"via-privato"`
	Civico                 int       `orm:"digits(4)" form:"" maxLength:"4" valid:"Required;MinSize(1);MaxSize(4);Numeric" id:"civico-privato"`
	Prefisso               string    `orm:"size(2)" form:"" maxLength:"2" valid:"Required;MaxSize(2);Numeric" id:"prefisso-privato"`
	Telefono               string    `orm:"size(10);unique" form:"" maxLength:"10" valid:"Required;Length(10);Tel" id:"telefono-privato"`
	Email                  string    `orm:"size(255);unique" form:"" valid:"Required;MaxSize(100);Email" id:"email-privato"`
	EmailConfermata        bool      `form:"-"`
	Psw                    string    `orm:"size(255)" form:"Password,password,Password: " valid:"Required;MinSize(8);" id:"password-privato"`
	ConfermaPsw            string    `orm:"size(255)" form:"ConfermaPassword,password,Conferma password: " valid:"Required;MinSize(8);" id:"conferma-password-privato"`
	DataNascita            time.Time `orm:"type(date)" valid:"Required;Numeric"`
	Medico                 *Medico   `orm:"rel(fk);null;on_delete(set_null)" form:"-"`
}

// TODO: criptare password
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
