package models

import "github.com/beego/beego/v2/client/orm"

func init() {
	orm.RegisterModel(new(InfoTest))
}

type InfoTest struct {
	IdInfoTest    int          `orm:"pk;auto"`
	IdLaboratorio *Laboratorio `orm:"rel(fk)"`
	TipologiaTest string
	Costo         float32
	Tempi         int // Tempo espressi in secondi
}
