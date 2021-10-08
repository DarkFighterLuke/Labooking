package utils

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
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
	qs := o.QueryTable("recupero_password")

	hour, minute, _ := time.Now().Clock()
	orario := strconv.Itoa(hour) + ":" + strconv.Itoa(minute) + ":" + "00"

	x, err := qs.Filter("timeout__lt", orario).Delete()
	fmt.Printf("[%v:%v:%v] ", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	fmt.Println("Timer expired hashcode over, number of deleted hash: ", x)
	if err != nil {
		return err
	} else {
		return nil
	}
}
