package models

import "github.com/beego/beego/v2/client/orm"

func init() {
	orm.RegisterModel(new(InfoTest))
}

type InfoTest struct {
	IdInfoTest    int64        `orm:"pk;auto"`
	IdLaboratorio *Laboratorio `orm:"rel(fk);column(id_laboratorio)"`
	TipologiaTest string       `orm:""`
	Costo         float64      `orm:"digits(6);decimals(2)"`
	Tempi         int64        `orm:""` // Tempo espresso in secondi
}

func (it *InfoTest) Aggiungi() (int64, error) {
	o := orm.NewOrm()
	newId, err := o.Insert(it)
	return newId, err
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
