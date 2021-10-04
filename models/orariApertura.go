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
