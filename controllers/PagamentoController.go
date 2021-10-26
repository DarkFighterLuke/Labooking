package controllers

import (
	"Labooking/models"
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

type PagamentoController struct{}

func (pc *PagamentoController) Paga(prenotazioneController web.Controller, td *models.TestDiagnostico) error {
	numeroCartaStr := prenotazioneController.GetString("numero-carta")
	scadenzaStr := prenotazioneController.GetString("scadenza")
	cvvStr := prenotazioneController.GetString("cvv")
	err := pc.verificaDatiPagamento(numeroCartaStr, scadenzaStr, cvvStr)
	if err != nil {
		td.Pagato = false
	} else {
		td.Pagato = true
	}
	return err
}

func (pc *PagamentoController) verificaDatiPagamento(numeroCarta, scadenza, cvv string) (err error) {
	if len(numeroCarta) != 16 {
		return fmt.Errorf("formato numero carta errato")
	}
	if len(scadenza) == 0 {
		return fmt.Errorf("formato scadenza carta errato")
	}
	if len(cvv) != 3 {
		return fmt.Errorf("formato cvv carta errato")
	}

	return nil
}
