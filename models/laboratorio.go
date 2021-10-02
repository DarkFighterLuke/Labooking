package models

import "github.com/beego/beego/v2/client/orm"

func init() {
	orm.RegisterModel(new(Laboratorio))
}

type Laboratorio struct {
	IdLaboratorio   int    `orm:"pk;auto" form:"-"`
	Nome            string `orm:"size(255)" form:""`
	PartitaIva      string `orm:"size(11);unique" form:",,Partita Iva: " maxLenght:"11"`
	Citta           string `orm:"size(255)" form:""`
	Cap             string `orm:"size(5)" form:",,CAP: " maxLength:"5"`
	Via             string `orm:"size(255)" form:",,Via/Piazza: "`
	Civico          int    `orm:"digits(4)" form:"" maxLength:"4"`
	Prefisso        string `orm:"size(2)" form:"" maxLength:"2"`
	Telefono        string `orm:"size(10);unique" form:"" maxLength:"10"`
	Email           string `orm:"size(255);unique" form:""`
	EmailConfermata bool   `form:"-"`
	Psw             string `orm:"size(255)" form:"Password,password,Password: "`
}

func (l *Laboratorio) Aggiungi() error {
	o := orm.NewOrm()
	_, err := o.Insert(l)
	return err
}

func (l *Laboratorio) Seleziona() (interface{}, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(l)
	laboratori := make([]Laboratorio, 0)
	_, err := qs.All(laboratori)
	if err != nil {
		return nil, err
	}
	return laboratori, nil
}

func (l *Laboratorio) Modifica() error {
	o := orm.NewOrm()
	_, err := o.Update(l)
	return err
}

func (l *Laboratorio) Elimina() error {
	o := orm.NewOrm()
	_, err := o.Delete(l)
	return err
}
