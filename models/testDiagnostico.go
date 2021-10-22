package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(TestDiagnostico))
}

type TestDiagnostico struct {
	IdTestDiagnostico int          `orm:"pk;auto"`
	DataPrenotazione  time.Time    `orm:"type(date)"`
	DataEsecuzione    time.Time    `orm:"type(date)"`
	Pagato            bool         `orm:""`
	TipologiaTest     string       `orm:""`
	Stato             string       `orm:""`
	Laboratorio       *Laboratorio `orm:"rel(fk);on_update(cascade);on_delete(do_nothing);column(id_laboratorio)"`
	Referto           *Referto     `orm:"rel(fk);on_update(cascade);on_delete(cascade);column(id_referto);null"`
	Privato           *Privato     `orm:"rel(fk);on_update(cascade);on_delete(cascade);column(id_privato)"`
}

func (td *TestDiagnostico) Aggiungi() (int64, error) {
	o := orm.NewOrm()
	newId, err := o.Insert(td)
	if err != nil {
		return newId, err
	}
	return newId, err
}

func (td *TestDiagnostico) Modifica() error {
	o := orm.NewOrm()
	_, err := o.Update(td)
	return err
}

func (td *TestDiagnostico) Elimina(cols ...string) error {
	o := orm.NewOrm()
	err := td.Seleziona(cols...)
	if err != nil {
		return err
	}
	_, err = o.Delete(td)
	if err != nil {
		return err
	}
	return nil
}

func (td *TestDiagnostico) Seleziona(cols ...string) error {
	o := orm.NewOrm()
	err := o.Read(td, cols...)
	if err != nil {
		return err
	}
	return nil
}

func (td *TestDiagnostico) SelezionaByDataStr(dataStr, slot string) error {
	o := orm.NewOrm()
	err := o.Raw("SELECT * FROM test_diagnostico WHERE id_laboratorio=? AND data_esecuzione=?", td.Laboratorio.IdLaboratorio, dataStr+" "+slot).QueryRow(td)
	return err
}
