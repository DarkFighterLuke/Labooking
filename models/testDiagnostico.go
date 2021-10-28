package models

import (
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"time"
)

func init() {
	orm.RegisterModel(new(TestDiagnostico))
}

type TestDiagnostico struct {
	IdTestDiagnostico int64                 `orm:"pk;auto"`
	DataPrenotazione  time.Time             `orm:"type(date)"`
	DataEsecuzione    time.Time             `orm:"type(date)"`
	Pagato            bool                  `orm:""`
	TipologiaTest     string                `orm:""`
	Stato             string                `orm:""`
	Laboratorio       *Laboratorio          `orm:"rel(fk);on_update(cascade);on_delete(do_nothing);column(id_laboratorio)"`
	Referto           *Referto              `orm:"rel(fk);on_update(cascade);on_delete(cascade);column(id_referto);null"`
	Privato           *Privato              `orm:"rel(fk);on_update(cascade);on_delete(cascade);column(id_privato)"`
	Questionario      *QuestionarioAnamnesi `orm:"reverse(one);column(id_questionario_anamnesi)"`
}

func (td *TestDiagnostico) Aggiungi() (int64, error) {
	o := orm.NewOrm()
	newId, err := o.Insert(td)
	if err != nil {
		return newId, err
	}
	return newId, err
}

func (td *TestDiagnostico) Modifica() error {
	o := orm.NewOrm()
	_, err := o.Update(td)
	return err
}

func (td *TestDiagnostico) Elimina(cols ...string) error {
	o := orm.NewOrm()
	err := td.Seleziona(cols...)
	if err != nil {
		return err
	}
	_, err = o.Delete(td)
	if err != nil {
		return err
	}
	return nil
}

func (td *TestDiagnostico) Seleziona(cols ...string) error {
	o := orm.NewOrm()
	err := o.Read(td, cols...)
	if err != nil {
		return err
	}
	return nil
}

func (td *TestDiagnostico) SelezionaByDataStr(dataStr, slot string) error {
	o := orm.NewOrm()
	err := o.Raw("SELECT * FROM test_diagnostico WHERE id_laboratorio=? AND data_esecuzione=?", td.Laboratorio.IdLaboratorio, dataStr+" "+slot).QueryRow(td)
	return err
}

func SelezionaTestAll() (testDiagnostici []*TestDiagnostico, err error) {
	o := orm.NewOrm()

	_, err = o.QueryTable("test_diagnostico").RelatedSel().All(&testDiagnostici)
	for _, v := range testDiagnostici {
		v.LoadRelatedQuestionari()
		v.LoadRelatedReferto()
	}
	return testDiagnostici, err
}

func (td *TestDiagnostico) LoadRelatedQuestionari() {
	qa := new(QuestionarioAnamnesi)
	qa.TestDiagnostico = new(TestDiagnostico)

	qa.TestDiagnostico.IdTestDiagnostico = td.IdTestDiagnostico
	err := qa.Seleziona("id_test_diagnostico")
	if err != nil {
		return
	}
	td.Questionario = qa
}

func (td *TestDiagnostico) LoadRelatedReferto() {
	r := new(Referto)
	r.IdReferto = td.IdTestDiagnostico
	err := r.Seleziona("id_referto")
	if err != nil {
		return
	}
	td.Referto = r
}

func (td *TestDiagnostico) CheckInviaMailiOrganizzazione() (bool, error) {
	o := orm.NewOrm()
	var prenotazioni int
	var prenotazioniNotificate int

	dataStr := td.DataPrenotazione.Format("2006-01-02")
	organizzazioneStr := strconv.Itoa(td.Privato.Organizzazione.IdOrganizzazione)
	queryPrenotazioni := "SELECT COUNT(*) FROM test_diagnostico td, privato p WHERE data_prenotazione = '" + dataStr + "' AND td.id_privato = p.id_privato AND p.organizzazione = '" + organizzazioneStr + "'"
	err := o.Raw(queryPrenotazioni).QueryRow(&prenotazioni)
	if err != nil {
		return false, err
	}

	queryPrenotazioniNotificate := "SELECT COUNT(*) FROM test_diagnostico td, privato p WHERE data_prenotazione = '" + dataStr + "' AND td.id_privato = p.id_privato AND p.organizzazione = '" + organizzazioneStr + "' AND stato ='notificato'"
	err = o.Raw(queryPrenotazioniNotificate).QueryRow(&prenotazioniNotificate)
	if err != nil {
		return false, err
	}

	if prenotazioni == prenotazioniNotificate {
		return true, nil
	} else {
		return false, nil
	}

}
func (td *TestDiagnostico) SelezionaAllTestsByPrivatoStato() (testDiagnostici []*TestDiagnostico, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("test_diagnostico").Filter("id_privato", td.Privato.IdPrivato).Filter("stato", td.Stato).All(&testDiagnostici)
	return testDiagnostici, err
}
