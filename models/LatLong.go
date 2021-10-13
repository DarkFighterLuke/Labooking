package models

import (
	"encoding/json"
	"net/http"
	url2 "net/url"
)

type LatLong struct {
	PlaceId int64  `json:"place_id"`
	License string `json:"licence"`
	Lat     string `json:"lat"`
	Long    string `json:"lon"`
}

func RetrieveLatLong(indirizzo string, ll *LatLong) error {
	url := "https://nominatim.openstreetmap.org/search?format=json&q=" + url2.QueryEscape(indirizzo)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	var llArray []LatLong

	jsonDecoder := json.NewDecoder(resp.Body)
	err = jsonDecoder.Decode(&llArray)

	if len(llArray) > 0 {
		*ll = llArray[0]
	}

	return err
}
