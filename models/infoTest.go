package models

import (
	"github.com/beego/beego/v2/client/orm"
)

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

func (it *InfoTest) SelezionaMinCosto() error {
	o := orm.NewOrm()
	query := "SELECT MIN(it.costo) AS costo FROM info_test it"
	err := o.Raw(query).QueryRow(&it)
	return err
}

func (it *InfoTest) SelezionaMaxCosto() error {
	o := orm.NewOrm()
	query := "SELECT MAX(it.costo) AS costo FROM info_test it"
	err := o.Raw(query).QueryRow(&it)
	return err
}

func (it *InfoTest) SelezionaMinTempi() error {
	o := orm.NewOrm()
	query := "SELECT MIN(it.tempi) AS tempi FROM info_test it"
	err := o.Raw(query).QueryRow(&it)
	return err
}

func (it *InfoTest) SelezionaMaxTempi() error {
	o := orm.NewOrm()
	query := "SELECT MAX(it.tempi) AS tempi FROM info_test it"
	err := o.Raw(query).QueryRow(&it)
	return err
}

func (it *InfoTest) SelezionaInfoTestByLabId() (its []*InfoTest, err error) {
	o := orm.NewOrm()
	query := "SELECT * FROM info_test WHERE id_laboratorio=?"
	_, err = o.Raw(query, it.IdLaboratorio.IdLaboratorio).QueryRows(&its)
	return its, err
}
