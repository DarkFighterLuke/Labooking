package models

import "github.com/beego/beego/v2/client/orm"

func init() {
	orm.RegisterModel(new(InfoTest))
}

type InfoTest struct {
	IdInfoTest    int          `orm:"pk;auto"`
	IdLaboratorio *Laboratorio `orm:"rel(fk)"`
	TipologiaTest string       `orm:""`
	Costo         float32      `orm:"digits(4,2)"`
	Tempi         int          `orm:""` // Tempo espresso in secondi
}
