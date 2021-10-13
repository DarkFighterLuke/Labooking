package utils

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

func Timer() {
	deleteHash := time.NewTicker(30 * time.Second)

	for {
		select {
		case <-deleteHash.C:
			err := deleteHashExpired()
			if err != nil {
				fmt.Println("Errore timer: ", err)
			}
		}
	}
}

func deleteHashExpired() error {
	o := orm.NewOrm()
	res, err := o.Raw("DELETE FROM recupero_password WHERE timeout < NOW();").Exec()
	fmt.Printf("[%v:%v:%v] ", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	x, _ := res.RowsAffected()
	fmt.Println("Timer expired hashcode over, number of deleted hash: ", x)
	if err != nil {
		return err
	}

	return nil
}
