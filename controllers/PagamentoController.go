package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
)

type PagamentoController struct{}

func (pc *PagamentoController) Paga(prenotazioneController web.Controller) error {
	numeroCartaStr := prenotazioneController.GetString("numero-carta")
	scadenzaStr := prenotazioneController.GetString("scadenza")
	cvvStr := prenotazioneController.GetString("cvv")
	err := pc.VerificaDatiPagamento(numeroCartaStr, scadenzaStr, cvvStr)
	return err
}

func (pc *PagamentoController) VerificaDatiPagamento(numeroCarta, scadenza, cvv string) (err error) {
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
