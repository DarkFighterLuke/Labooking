package models

import (
	"Labooking/models/utils"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(RecuperoPassword))
}

type GestioneRecupero interface {
	CreaRecupero() (string, error)
}

type RecuperoPassword struct {
	HashKey        string          `orm:"size(64);pk"`
	Timeout        time.Time       `orm:""`
	Medico         *Medico         `orm:"rel(one);null;on_delete(do_nothing);on_update(cascade);column(id_medico)"`
	Privato        *Privato        `orm:"rel(one);null;on_delete(do_nothing);on_update(cascade);column(id_privato)"`
	Laboratorio    *Laboratorio    `orm:"rel(one);null;on_delete(do_nothing);on_update(cascade);column(id_laboratorio)"`
	Organizzazione *Organizzazione `orm:"rel(one);null;on_delete(do_nothing);on_update(cascade);column(id_organizzazione)"`
}

func (m *Medico) CreaRecupero() (string, error) {
	err := m.Seleziona("email")
	if err != nil {
		return "", err
	}

	rec := generaHash()
	rec.Medico = new(Medico)
	rec.Medico.IdMedico = m.IdMedico

	err = insertOrUpdate(&rec, "id_medico")
	if err != nil {
		return "", err
	}

	return rec.HashKey, nil
}

func (p *Privato) CreaRecupero() (string, error) {
	err := p.Seleziona("email")
	if err != nil {
		return "", err
	}

	rec := generaHash()
	rec.Privato = new(Privato)
	rec.Privato.IdPrivato = p.IdPrivato

	err = insertOrUpdate(&rec, "id_privato")
	if err != nil {
		return "", err
	}

	return rec.HashKey, nil
}

func (l *Laboratorio) CreaRecupero() (string, error) {
	err := l.Seleziona("email")
	if err != nil {
		return "", err
	}

	rec := generaHash()
	rec.Laboratorio = new(Laboratorio)
	rec.Laboratorio.IdLaboratorio = l.IdLaboratorio

	err = insertOrUpdate(&rec, "id_laboratorio")
	if err != nil {
		return "", err
	}

	return rec.HashKey, nil
}

func (org *Organizzazione) CreaRecupero() (string, error) {
	err := org.Seleziona("email")
	if err != nil {
		return "", err
	}

	rec := generaHash()
	rec.Organizzazione = new(Organizzazione)
	rec.Organizzazione.IdOrganizzazione = org.IdOrganizzazione

	err = insertOrUpdate(&rec, "id_organizzazione")
	if err != nil {
		return "", err
	}

	return rec.HashKey, nil
}

func generaHash() RecuperoPassword {
	hashKey := utils.RandStringRunes(64)

	r := new(RecuperoPassword)
	r.HashKey = hashKey
	r.Timeout = time.Now().Local().Add(time.Hour)

	return *r
}

func (rp *RecuperoPassword) CercaHash() error {
	o := orm.NewOrm()

	err := o.Read(rp)
	if err != nil {
		return err
	}
	return nil
}

func (rp *RecuperoPassword) Elimina() error {
	o := orm.NewOrm()

	_, err := o.Delete(rp)
	if err != nil {
		return err
	}
	return nil
}

func insertOrUpdate(rec *RecuperoPassword, idCol string) error {
	o := orm.NewOrm()

	newRec := *rec
	err := o.Read(rec, idCol)
	if err != nil {
		_, err := o.Insert(rec)
		if err != nil {
			return err
		}
	} else {
		_, err := o.Delete(rec)
		if err != nil {
			return err
		}

		_, err = o.Insert(&newRec)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteHashExpired() error {
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM recupero_password WHERE timeout < NOW();").Exec()
	//fmt.Printf("[%v:%v:%v] ", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	//x, _ := res.RowsAffected()
	//fmt.Println("Timer expired hashcode over, number of deleted hash: ", x)
	if err != nil {
		return err
	}

	return nil
}
