package models

import "time"

type OrariApertura struct {
	IdOrariApertura int
	IdLaboratorio   int
	Giorno          int
	Orario          time.Time
	Stato           bool
}
