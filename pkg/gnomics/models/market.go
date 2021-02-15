package models

type Market struct {
	Exchange string `json:"exchange"`
	Market   string `json:"market"`
	Base     string `json:"base"`
	Quote    string `json:"quote"`
}
