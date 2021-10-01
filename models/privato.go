package models

import "time"

type Privato struct {
	IdPrivato              int    `orm:"pk;auto"`
	Nome                   string `orm:"size(255)"`
	Cognome                string `orm:"size(255)"`
	CodiceFiscale          string `orm:"size(16);unique"`
	NumeroTesseraSanitaria string `orm:"size(20);unique"`
	Citta                  string `orm:"size(255)"`
	Cap                    string `orm:"size(5)"`
	Via                    string `orm:"size(255)"`
	Civico                 int    `orm:"digits(4)"`
	Prefisso               string `orm:"size(2)"`
	Telefono               string `orm:"size(10);unique"`
	Email                  string `orm:"size(255);unique"`
	EmailConfermata        bool
	Psw                    string    `orm:"size(255)"`
	DataNascita            time.Time `orm:"type(date)"`
	Medico                 *Medico   `orm:"rel(fk)"`
}
