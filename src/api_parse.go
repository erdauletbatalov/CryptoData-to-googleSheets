package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var cryptorank CryptoRank
var coingecko CoinGecko

// Adds names of colomns and fills Spreadsheet with first 3 currenty from CryptoRank
func GetAPIFromCryptoRank(numberOfRows int) [][]interface{} {
	rows := make([][]interface{}, 0)
	currencies := GetRequest("https://api.cryptorank.io/v1/currencies?api_key=ce6f0d432f8a3326d198dcf6d99c0d745aecd0a26ed040450c7ee796b236")
	err := json.Unmarshal(currencies, &cryptorank)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	rows = append(rows, []interface{}{"Наименование", "Теги", "TIMESTAMP"})
	for i, val := range cryptorank.Data {
		if len(cryptorank.Data) < numberOfRows {
			log.Println(fmt.Errorf("Too many rows to read"))
			os.Exit(1)
		}
		if i == numberOfRows {
			break
		}
		rows = append(rows, []interface{}{val.Name, val.Category, time.Time.String(val.LastUpdated)})
	}
	return rows
}

func GetAPIFromCoinGecko(numberOfRows int) [][]interface{} {
	rows := make([][]interface{}, 0)
	currencies := GetRequest("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc")
	err := json.Unmarshal(currencies, &coingecko)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	rows = append(rows, []interface{}{"Наименование", "Цена", "TIMESTAMP"})
	for i, val := range coingecko {
		if len(cryptorank.Data) < numberOfRows {
			log.Println(fmt.Errorf("Too many rows to read"))
			os.Exit(1)
		}
		if i == numberOfRows {
			break
		}
		rows = append(rows, []interface{}{val.Name, val.CurrentPrice, time.Time.String(val.LastUpdated)})
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
