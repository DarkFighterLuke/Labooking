package models

import "github.com/beego/beego/v2/client/orm"

func init() {
	orm.RegisterModel(new(Medico))
}

type Medico struct {
	IdMedico        int    `orm:"pk;auto" form:"-"`
	Nome            string `orm:"size(255)" form:""`
	Cognome         string `orm:"size(255)" form:""`
	CodiceFiscale   string `orm:"size(16);unique" form:",,Codice Fiscale: " maxLength:"16"`
	CittaStudio     string `orm:"size(255)" form:",,Città dello studio: "`
	CapStudio       string `orm:"size(5)" form:",,CAP dello studio: " maxLength:"5"`
	ViaStudio       string `orm:"size(255)" form:",,Via/Piazza dello studio: "`
	CivicoStudio    int    `orm:"digits(4)" form:",,Civico dello studio: " maxLength:"4"`
	Prefisso        string `orm:"size(2)" form:"" maxLength:"2"`
	Telefono        string `orm:"size(10);unique" form:"" maxLength:"10"`
	Email           string `orm:"size(255);unique" form:""`
	EmailConfermata bool   `form:"-"`
	Psw             string `orm:"size(255)" form:"Password,password,Password: "`
	CodiceRegionale string `orm:"size(255)" form:",,Codice regionale: "`
}

func (m *Medico) Aggiungi() error {
	o := orm.NewOrm()
	_, err := o.Insert(m)
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
