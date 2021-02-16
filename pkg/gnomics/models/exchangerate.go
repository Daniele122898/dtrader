package models

type ExchangeRate struct {
	Currency  string `json:"currency"`
	Rate      string `json:"rate"`
	Timestamp string `json:"timestamp"`
}

type ExchangeRateHistory struct {
	Timestamp string `json:"timestamp"`
	Rate      string `json:"rate"`
}


