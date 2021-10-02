package models

import "time"

type OrariApertura struct {
	IdOrariApertura int          `orm:"pk;auto"`
	IdLaboratorio   *Laboratorio `orm:"rel(fk)"`
	Giorno          int
	Orario          time.Time
	Stato           bool
}
