package test

import (
	"Labooking/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestRetrieveLatLong(t *testing.T) {
	indirizzo := "Bari Via Sparano, 15"
	ll := models.LatLong{}
	err := models.RetrieveLatLong(indirizzo, &ll)
	if err != nil {
		t.Errorf("err is %v; want nil", err)
	}

	/*if ll.Lat==0 || ll.Long==0{
		t.Errorf("err is %v; want nil", err)
	}*/

	fmt.Printf("%#v", ll)
}
