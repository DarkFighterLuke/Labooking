package test

import (
	"Labooking/models"
	"testing"
	"time"
)

func TestAggiungi(t *testing.T) {
	p := new(models.Privato)

	p.Nome = "Tizio"
	p.Cognome = "Caio"
	p.CodiceFiscale = "123456789qwertyu"
	p.NumeroTesseraSanitaria = "12345678912345678911"
	p.Citta = "Roma"
	p.Cap = "00000"
	p.Via = "RRRRR"
	p.Civico = 123
	p.Prefisso = "39"
	p.Telefono = "234234234"
	p.Email = "tizio.caio@vvv.it"
	p.EmailConfermata = false
	p.Psw = "qwerty"
	p.DataNascita = time.Date(2000, 4, 3, 12, 30, 0, 0, time.UTC)
	p.Medico = nil

	err := p.Aggiungi()

	if err != nil {
		t.Errorf("err is %v; want nil", err)
	}
}
