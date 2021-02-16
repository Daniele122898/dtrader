package models

type Market struct {
	Exchange string `json:"exchange"`
	Market   string `json:"market"`
	Base     string `json:"base"`
	Quote    string `json:"quote"`
}

type MarketCap struct {
	Timestamp            string `json:"timestamp"`
	MarketCap            string `json:"market_cap"`
	TransparentMarketCap string `json:"transparent_market_cap"`
}

