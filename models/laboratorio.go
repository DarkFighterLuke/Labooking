package models

import (
	"Labooking/models/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
)

func init() {
	orm.RegisterModel(new(Laboratorio))
}

type Laboratorio struct {
	IdLaboratorio int64            `orm:"pk;auto" form:"-" json:"id_laboratorio"`
	Nome          string           `orm:"size(255)" form:"" valid:"Required" id:"nome-laboratorio" json:"nome"`
	PartitaIva    string           `orm:"size(11);unique" form:",,Partita Iva: " maxLength:"11" valid:"Required;Length(11)" id:"partita-iva-laboratorio"`
	Indirizzo     string           `orm:"size(255)" form:"" valid:"Required" id:"indirizzo-laboratorio"`
	Lat           float64          `orm:"digits(10);decimals(7)" form:"-" json:"lat"`
	Long          float64          `orm:"digits(10);decimals(7)" form:"-" json:"long"`
	Prefisso      string           `orm:"size(6)" form:"-" valid:"Required"`
	Telefono      string           `orm:"size(10);unique" form:"" maxLength:"10" valid:"Required;Numeric;Length(10)" id:"telefono-laboratorio"`
	Email         string           `orm:"size(255);unique" form:"" valid:"Required;Email" id:"email-laboratorio"`
	Psw           string           `orm:"size(255)" form:"Password,password,Password: " valid:"Required" id:"password-laboratorio"`
	ConfermaPsw   string           `orm:"-" form:"ConfermaPassword,password,Conferma password: " valid:"Required" id:"conferma-password-laboratorio"`
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

func FiltraLaboratori(laboratori *[]Laboratorio, tempo int64, tipi map[string]bool, costo float64, orario_inizio string, orario_fine string, giorno string) error {
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
	if giorno != "" {
		giornoQuery = "AND oa.giorno = '" + giorno + "'"
	}

	var orarioAperturaQuery string
	if orario_inizio != "" {
		orarioAperturaQuery = "AND l.id_laboratorio IN (SELECT oa.id_laboratorio " +
			"FROM orari_apertura oa " +
			"WHERE oa.orario >= " + orario_inizio + " AND oa.stato = 1 AND l.id_laboratorio = oa.id_laboratorio " + giornoQuery + ") "
	}

	var orarioChiusuraQuery string
	if orario_fine != "" {
		orarioChiusuraQuery = "AND l.id_laboratorio IN (SELECT oa.id_laboratorio " +
			"FROM orari_apertura oa " +
			"WHERE oa.orario <= " + orario_fine + " AND oa.stato = 0 AND l.id_laboratorio = oa.id_laboratorio " + giornoQuery + " ) "
	}

	query := "SELECT l.id_laboratorio, l.nome, l.lat, l.long " +
		"FROM laboratorio l, info_test it " +
		"WHERE l.id_laboratorio = it.id_laboratorio AND " + tipiQuery + " it.tempi <= ? AND it.costo <= ? " +
		orarioAperturaQuery + orarioChiusuraQuery

	if orarioAperturaQuery == "" && orarioChiusuraQuery == "" && giorno != "" {
		query = query + "AND l.id_laboratorio IN (SELECT oa.id_laboratorio " +
			"FROM orari_apertura oa " +
			"WHERE oa.giorno = '" + giorno + "') "
	}

	_, err := o.Raw(query, tempo, costo).QueryRows(laboratori)
	return err
}

func PrelevaLaboratoriForMap(laboratori *[]Laboratorio) error {
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
