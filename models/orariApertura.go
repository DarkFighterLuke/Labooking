package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(OrariApertura))
}

type OrariApertura struct {
	IdOrariApertura int64        `orm:"pk;auto"`
	IdLaboratorio   *Laboratorio `orm:"rel(fk);column(id_laboratorio)"`
	Giorno          string       `orm:""`
	Orario          time.Time    `orm:"type(date)"`
	Stato           bool         `orm:"type(bool)"`
}

func (oa *OrariApertura) Aggiungi() (int64, error) {
	o := orm.NewOrm()
	newId, err := o.Insert(oa)
	return newId, err
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

func (oa *OrariApertura) SelezionaOrariAperturaByIdLab() (oas []*OrariApertura, err error) {
	o := orm.NewOrm()
	query := "SELECT * FROM orari_apertura oas WHERE oas.stato=1 AND oas.id_laboratorio=? ORDER BY oas.id_orari_apertura"
	_, err = o.Raw(query, oa.IdLaboratorio).QueryRows(&oas)
	return oas, err
}

func (oa *OrariApertura) SelezionaOrariChiusuraByIdLab() (oas []*OrariApertura, err error) {
	o := orm.NewOrm()
	query := "SELECT * FROM orari_apertura oas WHERE oas.stato=0 AND oas.id_laboratorio=? ORDER BY oas.id_orari_apertura"
	_, err = o.Raw(query, oa.IdLaboratorio).QueryRows(&oas)
	return oas, err
}
