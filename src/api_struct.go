package service

import "time"

type CryptoRank struct {
	Data []struct {
		Name        string    `json:"name"`
		Category    string    `json:"category"`
		LastUpdated time.Time `json:"lastUpdated"`
	} `json:"data"`
}

type CoinGecko []struct {
	Name         string    `json:"name"`
	CurrentPrice float32   `json:"current_price"`
	LastUpdated  time.Time `json:"last_updated"`
}
