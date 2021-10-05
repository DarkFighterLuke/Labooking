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
	IdLaboratorio   *Laboratorio `orm:"rel(fk)"`
	Giorno          int
	Orario          time.Time
	Stato           bool
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
