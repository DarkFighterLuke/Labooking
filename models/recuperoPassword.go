package models

import (
	"Labooking/models/utils"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"net/smtp"
	"time"
)

func init() {
	orm.RegisterModel(new(RecuperoPassword))
}

type GestioneRecupero interface {
	creaRecupero() (string, error)
	InviaLink() error
}

type RecuperoPassword struct {
	HashKey        string          `orm:"size(64);pk"`
	Timeout        time.Time       `orm:""`
	Medico         *Medico         `orm:"rel(one);null;on_delete(do_nothing);on_update(cascade);column(id_medico)"`
	Oganizzazione  *Organizzazione `orm:"rel(one);on_delete(do_nothing);on_update(cascade);column(id_organizzazione)"`
	Privato        *Privato        `orm:"rel(one);null;on_delete(do_nothing);on_update(cascade);column(id_privato)"`
	Laboratorio    *Laboratorio    `orm:"rel(one);null;on_delete(do_nothing);on_update(cascade);column(id_laboratorio)"`
	Organizzazione *Organizzazione `orm:"rel(one);null;on_delete(do_nothing);on_update(cascade);column(id_organizzazione)"`
}

func (m *Medico) creaRecupero() (string, error) {
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

func (m *Medico) InviaLink() error {
	hashKey, err := m.creaRecupero()
	if err != nil {
		return err
	}

	err = inviaMail(hashKey, m.Email)
	if err != nil {
		return err
	}

	return nil
}

func (p *Privato) creaRecupero() (string, error) {
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

func (p *Privato) InviaLink() error {
	hashKey, err := p.creaRecupero()
	if err != nil {
		return err
	}

	err = inviaMail(hashKey, p.Email)
	if err != nil {
		return err
	}

	return nil
}

func (l *Laboratorio) creaRecupero() (string, error) {
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

func (l *Laboratorio) InviaLink() error {
	hashKey, err := l.creaRecupero()
	if err != nil {
		return err
	}

	err = inviaMail(hashKey, l.Email)
	if err != nil {
		return err
	}

	return nil
}

func (org *Organizzazione) creaRecupero() (string, error) {
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

func (org *Organizzazione) InviaLink() error {
	hashKey, err := org.creaRecupero()
	if err != nil {
		return err
	}

	err = inviaMail(hashKey, org.Email)
	if err != nil {
		return err
	}

	return nil
}

func inviaMail(hashKey, reciver string) error {
	from, err := web.AppConfig.String("email")
	if err != nil {
		return err
	}
	pass, err := web.AppConfig.String("passwordemail")
	if err != nil {
		return err
	}
	to := reciver

	websitelink, err := web.AppConfig.String("websitelink")
	if err != nil {
		return err
	}

	link := websitelink + "cambiapassword?hash=" + hashKey

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Recupero password Labooking\n\n" +
		"Cliccare il seguente link per recuperare la password:\n" + link

	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		return err
	}
	return nil
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
