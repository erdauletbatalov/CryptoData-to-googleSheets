package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var Currencies Currency

func GetAPIFromCryptoRank() {
	currencies := GetRequest("https://api.cryptorank.io/v1/currencies?api_key=ce6f0d432f8a3326d198dcf6d99c0d745aecd0a26ed040450c7ee796b236")
	err := json.Unmarshal(currencies, &Currencies)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for i, val := range Currencies.Data {
		if i == 3 {
			return
		}
		fmt.Println(val)
	}
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
