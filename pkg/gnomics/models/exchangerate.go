package models

type ExchangeRate struct {
	Currency  string `json:"currency"`
	Rate      string `json:"rate"`
	Timestamp string `json:"timestamp"`
}

