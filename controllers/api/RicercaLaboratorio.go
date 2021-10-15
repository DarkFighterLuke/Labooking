package api

import (
	"Labooking/models"
	"github.com/beego/beego/v2/server/web"
	"strconv"
	"time"
)

type RicercaLaboratorio struct {
	web.Controller
}

func (rl *RicercaLaboratorio) Get() {
	tempo := rl.GetString("tempo")
	tempoInt, err := strconv.Atoi(tempo)
	if err != nil {
		rl.Ctx.WriteString("ricerca: " + err.Error())
		return
	}
	costo := rl.GetString("tipo")
	costoFloat, err := strconv.ParseFloat(costo, 64)
	orarioInizio := rl.GetString("inizio-intervallo")
	orarioFine := rl.GetString("fine-intervallo")
	giorno := rl.GetString("giorno")
	orarioIniziotime, err := time.Parse("15:04", orarioInizio)
	if err != nil {
		rl.Ctx.WriteString("ricerca: " + err.Error())
		return
	}
	orarioFinetime, err := time.Parse("15:04", orarioFine)
	if err != nil {
		rl.Ctx.WriteString("ricerca: " + err.Error())
		return
	}
	tipi := []string{rl.GetString("molecolare"), rl.GetString("antigenico"), rl.GetString("sierologico")}

	var labs []models.Laboratorio
	err = models.FiltraLaboratori(&labs, int64(tempoInt), tipi, costoFloat, orarioIniziotime, orarioFinetime, giorno)
	if err != nil {
		rl.Ctx.WriteString("ricerca: " + err.Error())
		return
	}

	rl.Data["json"] = &labs
	err = rl.ServeJSON()
	if err != nil {
		rl.Ctx.WriteString("ricerca: " + err.Error())
		return
	}
}
