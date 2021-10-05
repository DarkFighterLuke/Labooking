package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(OrariApertura))
}

type OrariApertura struct {
	IdOrariApertura int          `orm:"pk;auto"`
	IdLaboratorio   *Laboratorio `orm:"rel(fk);column(id_laboratorio)"`
	Giorno          string       `orm:""`
	Orario          time.Time    `orm:"type(date)"`
	Stato           bool         `orm:"type(bool)"`
}

func (oa *OrariApertura) Aggiungi() error {
	o := orm.NewOrm()
	_, err := o.Insert(oa)
	return err
}

func (oa *OrariApertura) Seleziona(cols ...string) error {
	o := orm.NewOrm()
	err := o.Read(oa, cols...)
	if err != nil {
		return err
	}
	return nil
}

func (oa *OrariApertura) Modifica() error {
	o := orm.NewOrm()
	_, err := o.Update(oa)
	return err
}

func (oa *OrariApertura) Elimina(cols ...string) error {
	o := orm.NewOrm()
	err := oa.Seleziona(cols...)
	_, err = o.Delete(oa)
	return err
}
