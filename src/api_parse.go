package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var Currencies Currency

// Adds names of colomns and fills Spreadsheet with first 3 currenty from CryptoRank
func GetAPIFromCryptoRank() [][]interface{} {
	rows := make([][]interface{}, 0)
	currencies := GetRequest("https://api.cryptorank.io/v1/currencies?api_key=ce6f0d432f8a3326d198dcf6d99c0d745aecd0a26ed040450c7ee796b236")
	err := json.Unmarshal(currencies, &Currencies)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	rows = append(rows, []interface{}{"Наименование", "Теги", "TIMESTAMP"})
	for i, val := range Currencies.Data {
		if i == 3 {
			break
		}
		rows = append(rows, []interface{}{val.Name, val.Category, time.Time.String(val.LastUpdated)})
	}
	return rows
}

func GetRequest(url string) []byte {
	r, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}
