package models

import "github.com/beego/beego/v2/client/orm"

func init() {
	orm.RegisterModel(new(InfoTest))
}

type InfoTest struct {
	IdInfoTest    int          `orm:"pk;auto"`
	IdLaboratorio *Laboratorio `orm:"rel(fk)"`
	TipologiaTest string       `orm:""`
	Costo         float32      `orm:"digits(4);decimals(2)"`
	Tempi         int          `orm:""` // Tempo espresso in secondi
}

func (it *InfoTest) Aggiungi() error {
	o := orm.NewOrm()
	_, err := o.Insert(it)
	return err
}

func (it *InfoTest) Seleziona(cols ...string) error {
	o := orm.NewOrm()
	err := o.Read(it, cols...)
	if err != nil {
		return err
	}
	return nil
}

func (it *InfoTest) Modifica() error {
	o := orm.NewOrm()
	_, err := o.Update(it)
	return err
}

func (it *InfoTest) Elimina(cols ...string) error {
	o := orm.NewOrm()
	err := it.Seleziona(cols...)
	_, err = o.Delete(it)
	return err
}
