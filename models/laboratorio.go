package models

import (
	"Labooking/models/utils"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	"time"
)

func init() {
	orm.RegisterModel(new(Laboratorio))
}

type Laboratorio struct {
	IdLaboratorio int64            `orm:"pk;auto" form:"-" json:"id_laboratorio"`
	Nome          string           `orm:"size(255)" form:"" valid:"Required" id:"nome-laboratorio" json:"nome"`
	PartitaIva    string           `orm:"size(11);unique" form:",,Partita Iva: " maxLength:"11" valid:"Required;Length(11)" id:"partita-iva-laboratorio"`
	Iban          string           `orm:"size(30)" form:"" maxLength:"30" valid:"Required;Length(30)" id:"iban-laboratorio"`
	Indirizzo     string           `orm:"size(255)" form:"" valid:"Required" id:"indirizzo-laboratorio"`
	Lat           float64          `orm:"digits(10);decimals(7)" form:"-" json:"lat"`
	Long          float64          `orm:"digits(10);decimals(7)" form:"-" json:"long"`
	Prefisso      string           `orm:"size(6)" form:"-" valid:"Required"`
	Telefono      string           `orm:"size(10);unique" form:"" maxLength:"10" valid:"Required;Numeric;Length(10)" id:"telefono-laboratorio"`
	Email         string           `orm:"size(255);unique" form:"" valid:"Required;Email" id:"email-laboratorio"`
	Psw           string           `orm:"size(255)" form:"Password,password,Password: " valid:"Required" id:"password-laboratorio"`
	ConfermaPsw   string           `orm:"-" form:"ConfermaPassword,password,Conferma password: " valid:"Required" id:"conferma-password-laboratorio"`
	TestPerOra    int64            `orm:"column(test_per_ora);type(int)" valid:"Required" id:"test-per-ora-laboratorio"`
	Test          []*InfoTest      `orm:"reverse(many)" form:"-"`
	Orari         []*OrariApertura `orm:"reverse(many)" form:"-"`
}

func (l *Laboratorio) Aggiungi() (int64, error) {
	var err error
	l.Psw, err = utils.CryptSHA1(l.Psw)
	if err != nil {
		return -1, err
	}

	o := orm.NewOrm()
	newId, err := o.Insert(l)
	return newId, err
}

func (l *Laboratorio) Seleziona(cols ...string) error {
	o := orm.NewOrm()
	err := o.Read(l, cols...)
	if err != nil {
		return err
	}
	return nil
}

func (l *Laboratorio) Modifica() error {
	o := orm.NewOrm()
	_, err := o.Update(l)
	return err
}

func (l *Laboratorio) Elimina(cols ...string) error {
	o := orm.NewOrm()
	err := l.Seleziona(cols...)
	_, err = o.Delete(l)
	return err
}

func FiltraLaboratori(tempo int64, tipi map[string]bool, costo float64, orarioInizioStr string, orarioFineStr string, dataStr string, numPersone int) (laboratori []Laboratorio, err error) {
	o := orm.NewOrm()
	tipiQuery := ""
	giornoQuery := ""
	tipiCounter := 0
	i := 0
	for _, val := range tipi {
		if val {
			tipiCounter++
		}
	}
	for key, val := range tipi {
		if val {
			if i == 0 {
				tipiQuery = "(it.tipologia_test = '"
			} else {
				tipiQuery = tipiQuery + " OR it.tipologia_test = '"
			}

			if i == (tipiCounter - 1) {
				tipiQuery = tipiQuery + key + "') AND"
			} else {
				tipiQuery = tipiQuery + key + "'"
			}

			i++
		}
	}

	var data time.Time
	var giorno string
	if dataStr != "" {
		var err error
		data, err = time.Parse("2006-01-02", dataStr)
		if err != nil {
			return nil, err
		}
		giorno = parseDayOfWeek(data.Weekday())
	}
	if giorno != "" {
		giornoQuery = "AND oa.giorno = '" + giorno + "'"
	}

	var orarioAperturaQuery string
	if orarioInizioStr != "" {
		orarioAperturaQuery = "AND l.id_laboratorio IN (SELECT oa.id_laboratorio " +
			"FROM orari_apertura oa " +
			"WHERE oa.orario <= '" + orarioInizioStr + "' AND oa.stato = 1 AND l.id_laboratorio = oa.id_laboratorio " + giornoQuery + ") "
	}

	var orarioChiusuraQuery string
	if orarioFineStr != "" {
		orarioChiusuraQuery = "AND l.id_laboratorio IN (SELECT oa.id_laboratorio " +
			"FROM orari_apertura oa " +
			"WHERE oa.orario >= '" + orarioFineStr + "' AND oa.stato = 0 AND l.id_laboratorio = oa.id_laboratorio " + giornoQuery + " ) "
	}

	query := "SELECT l.id_laboratorio, l.test_per_ora, l.nome, l.lat, l.long " +
		"FROM laboratorio l, info_test it " +
		"WHERE l.id_laboratorio = it.id_laboratorio AND " + tipiQuery + " it.tempi <= ? AND it.costo <= ? " +
		orarioAperturaQuery + orarioChiusuraQuery

	if orarioAperturaQuery == "" && orarioChiusuraQuery == "" && giorno != "" {
		query = query + "AND l.id_laboratorio IN (SELECT oa.id_laboratorio " +
			"FROM orari_apertura oa " +
			"WHERE oa.giorno = '" + giorno + "') "
	}

	_, err = o.Raw(query, tempo, costo).QueryRows(&laboratori)
	if err != nil {
		return nil, err
	}

	if orarioInizioStr != "" && orarioFineStr != "" && dataStr != "" {
		var labDisponibili []Laboratorio // laboratori con slot prenotabili nell'intervallo specificato dall'utente
		for _, l := range laboratori {
			isDisponibile, _, _, err := VerificaSlotDisponibili(l, orarioInizioStr, orarioFineStr, dataStr, numPersone)
			if err != nil {
				return nil, err
			}
			if isDisponibile {
				labDisponibili = append(labDisponibili, l)
			}
		}
		laboratori = labDisponibili
	}

	return laboratori, err
}

func VerificaSlotDisponibili(l Laboratorio, orarioInizioStr, orarioFineStr, dataStr string, numPersone int) (isDisponibile bool, userSlots, slotsPrenotati []*time.Time, err error) {
	if l.TestPerOra < 1 {
		return false, nil, nil, fmt.Errorf("testPerOra non può essere inferiore ad 1")
	}

	var data time.Time
	var giorno string
	if dataStr != "" {
		var err error
		data, err = time.Parse("2006-01-02", dataStr)
		if err != nil {
			return false, nil, nil, err
		}
		giorno = parseDayOfWeek(data.Weekday())
	}

	var oa OrariApertura
	oa.IdLaboratorio = &l
	orari_apertura, err := oa.SelezionaOrariAperturaByIdLab()
	if err != nil {
		return false, nil, nil, err
	}
	orari_chiusura, err := oa.SelezionaOrariChiusuraByIdLab()
	if err != nil {
		return false, nil, nil, err
	}

	orarioInizio, err := time.ParseInLocation("15:04", orarioInizioStr, time.Local)
	if err != nil {
		return false, nil, nil, err
	}
	orarioFine, err := time.ParseInLocation("15:04", orarioFineStr, time.Local)
	if err != nil {
		return false, nil, nil, err
	}
	for i, _ := range orari_apertura {
		fmt.Println(orari_apertura[i].Orario.Local(), " apertura")
		fmt.Println(orarioInizio.Local())
		fmt.Println(orari_chiusura[i].Orario.Local(), " chiusura")
		fmt.Println(orarioFine.Local())
		if orari_apertura[i].Orario.Before(orarioInizio) && orari_chiusura[i].Orario.After(orarioFine) && orari_apertura[i].Giorno == giorno && orari_chiusura[i].Giorno == giorno {
			intervalloInMinuti := orari_chiusura[i].Orario.Sub(orari_apertura[i].Orario).Minutes()
			durataSlot := 60 / l.TestPerOra
			numSlot := l.TestPerOra * int64(intervalloInMinuti) / 60
			slots := make([]*time.Time, 1) // ogni slot contiene il suo orario di inizio
			ou := orari_apertura[i].Orario.Local()
			slots[0] = &ou
			for j := 1; j < int(numSlot); j++ {
				prevSlot := slots[j-1].Add(time.Duration(durataSlot) * time.Minute).Local()
				slots = append(slots, &prevSlot)
			}
			var startSlotIndex int
			var endSlotIndex int
			// cerca l'indice del primo slot all'interno dell'intervallo specificato dell'utente
			for j, s := range slots {
				if s.After(orarioInizio.Local()) {
					startSlotIndex = j
					break
				}
			}
			// cerca l'indice dell'ultimo slot all'interno dell'intervallo specificato dell'utente
			for j, s := range slots[startSlotIndex:] {
				if s.After(orarioFine) {
					endSlotIndex = j - 1
					break
				}
			}
			userSlots = slots[startSlotIndex : startSlotIndex+endSlotIndex] // sub slice contenente solo gli slot dell'intervallo specificato dall'utente

			var td TestDiagnostico
			var slotsPrenotati []*time.Time
			for _, us := range userSlots {
				td.IdTestDiagnostico = 0
				td.Laboratorio = &l
				td.DataPrenotazione = data.Add(time.Duration(us.Hour())*time.Hour + time.Duration(us.Minute())*time.Minute) //?
				_ = td.Seleziona("id_laboratorio", "data_prenotazione")
				if td.IdTestDiagnostico != 0 {
					slotsPrenotati = append(slotsPrenotati, &td.DataPrenotazione)
				}
			}
			if len(userSlots)-len(slotsPrenotati) >= numPersone {
				return true, userSlots, slotsPrenotati, nil
			}
		}
	}
	return false, nil, nil, err
}

func GetLaboratoriForMap(laboratori *[]Laboratorio) error {
	o := orm.NewOrm()
	query := "SELECT l.id_laboratorio, l.nome, l.lat, l.long FROM laboratorio l"
	_, err := o.Raw(query).QueryRows(laboratori)
	return err
}

func (l *Laboratorio) Valid(v *validation.Validation) {

	if l.Psw != l.ConfermaPsw {
		v.SetError("ConfermaPsw", "must be equal to Psw")
	}

	/*if !utils.IsPswValid(l.Psw) {
		v.SetError("Psw", "is not strong enough")
	}*/
}

func parseDayOfWeek(number time.Weekday) string {
	switch number {
	case 1:
		return "lunedi"
	case 2:
		return "martedi"
	case 3:
		return "mercoledi"
	case 4:
		return "giovedi"
	case 5:
		return "venerdi"
	case 6:
		return "sabato"
	case 0:
		return "domenica"
	default:
		return ""
	}
}
