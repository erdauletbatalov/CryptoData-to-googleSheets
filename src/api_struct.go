package service

import "time"

type Currency struct {
	Data []struct {
		Name        string    `json:"name"`
		Category    string    `json:"category"`
		LastUpdated time.Time `json:"lastUpdated"`
	} `json:"data"`
	Status struct {
		Time time.Time `json:"time"`
	} `json:"status"`
}
