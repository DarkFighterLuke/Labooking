package models

import "time"

type Privato struct {
	IdPrivato              int
	Nome                   string
	Cognome                string
	CodiceFiscale          string
	NumeroTesseraSanitaria string
	Citta                  string
	Cap                    string
	Via                    string
	Civico                 int
	Prefisso               string
	Telefono               string
	Email                  string
	EmailConfermata        bool
	Psw                    string
	DataNascita            time.Time
	// TODO: Aggiungere puntatore a medico
}
