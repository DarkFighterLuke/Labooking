package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(Referto))
}

type Referto struct {
	IdReferto    int       `orm:"pk;auto"`
	DataRilascio time.Time `orm:"type(date)"`
	Risultato    string    `orm:""`
}

func (r *Referto) Aggiungi() (int64, error) {
	o := orm.NewOrm()
	newId, err := o.Insert(r)
	if err != nil {
		return newId, err
	}
	return newId, err
}

func (r *Referto) Seleziona(cols ...string) error {
	o := orm.NewOrm()
	err := o.Read(r, cols...)
	if err != nil {
		return err
	}
	return nil
}

func (r *Referto) Elimina(cols ...string) error {
	o := orm.NewOrm()
	err := r.Seleziona(cols...)
	if err != nil {
		return err
	}
	_, err = o.Delete(r)
	if err != nil {
		return err
	}
	return nil
}

func (r *Referto) Modifica() error {
	o := orm.NewOrm()
	_, err := o.Update(r)
	return err
}
