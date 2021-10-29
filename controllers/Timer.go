package controllers

import (
	"Labooking/models"
	"fmt"
	"time"
)

func Timer() {
	deleteHash := time.NewTicker(30 * time.Second)

	for {
		select {
		case <-deleteHash.C:
			err := models.DeleteHashExpired()
			if err != nil {
				fmt.Println("Errore timer: ", err)
			}
		}
	}
}
