package models

import "github.com/beego/beego/v2/client/orm"

func init() {
	orm.RegisterModel(new(QuestionarioAnamnesi))
}

type QuestionarioAnamnesi struct {
	IdQuestionarioAnamnesi int              `orm:"pk;auto"`
	Nome                   string           `orm:"size(32)"`
	TestDiagnostico        *TestDiagnostico `orm:"rel(fk);on_update(cascade);on_delete(cascade);column(id_test_diagnostico)"`
}

func (qa *QuestionarioAnamnesi) Aggiungi() (int64, error) {
	o := orm.NewOrm()
	newId, err := o.Insert(qa)
	if err != nil {
		return newId, err
	}
	return newId, err
}
