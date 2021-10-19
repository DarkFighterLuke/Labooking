package api

import (
	"Labooking/models"
	"github.com/beego/beego/v2/server/web"
	"strconv"
)

type RicercaLaboratorio struct {
	web.Controller
}

func (rl *RicercaLaboratorio) Get() {
	var labs []models.Laboratorio
	err := models.GetLaboratoriForMap(&labs)
	if err != nil {
		rl.Ctx.WriteString("ricerca: " + err.Error())
		return
	}
	rl.Data["json"] = &labs
	err = rl.ServeJSON()
}

func (rl *RicercaLaboratorio) Post() {
	numPersone, err := rl.GetInt("numero-persone")
	if err != nil || numPersone < 1 {
		numPersone = 1
	}
	tempo := rl.GetString("tempo")
	tempoInt, err := strconv.Atoi(tempo)
	if err != nil {
		rl.Ctx.WriteString("ricerca: " + err.Error())
		return
	}
	tempoSeconds := tempoInt * 3600

	costo := rl.GetString("costo")
	costoFloat, err := strconv.ParseFloat(costo, 64)
	orarioInizio := rl.GetString("inizio-intervallo")
	orarioFine := rl.GetString("fine-intervallo")
	data := rl.GetString("data")

	tipi := make(map[string]bool)
	tipi["molecolare"], _ = rl.GetBool("molecolare")
	tipi["antigenico"], _ = rl.GetBool("antigenico")
	tipi["sierologico"], _ = rl.GetBool("sierologico")

	var labs []models.Laboratorio
	err = models.FiltraLaboratori(&labs, int64(tempoSeconds), tipi, costoFloat, orarioInizio, orarioFine, data, numPersone)
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
