package gnomics

import "encoding/json"

func UnmarshalCurrencyTicker(data []byte) ([]CurrencyTicker, error) {
	var r []CurrencyTicker
	err := json.Unmarshal(data, &r)
	return r, err
}

type CurrencyTicker struct {
	ID                string   `json:"id"`
	Currency          string   `json:"currency"`
	Symbol            string   `json:"symbol"`
	Name              string   `json:"name"`
	LogoURL           string   `json:"logo_url"`
	Status            string   `json:"status"`
	Price             string   `json:"price"`
	PriceDate         string   `json:"price_date"`
	PriceTimestamp    string   `json:"price_timestamp"`
	CirculatingSupply string   `json:"circulating_supply"`
	MaxSupply         string   `json:"max_supply"`
	MarketCap         string   `json:"market_cap"`
	NumExchanges      string   `json:"num_exchanges"`
	NumPairs          string   `json:"num_pairs"`
	NumPairsUnmapped  string   `json:"num_pairs_unmapped"`
	FirstCandle       string   `json:"first_candle"`
	FirstTrade        string   `json:"first_trade"`
	FirstOrderBook    string   `json:"first_order_book"`
	Rank              string   `json:"rank"`
	High              string   `json:"high"`
	HighTimestamp     string   `json:"high_timestamp"`
	I1H               Interval `json:"1h"`
	I1D               Interval `json:"1d"`
	I7D               Interval `json:"7d"`
	I30D              Interval `json:"30d"`
	I365D             Interval `json:"365d"`
	Ytd               Interval `json:"ytd"`
}

type Interval struct {
	Volume             string `json:"volume"`
	PriceChange        string `json:"price_change"`
	PriceChangePct     string `json:"price_change_pct"`
	VolumeChange       string `json:"volume_change"`
	VolumeChangePct    string `json:"volume_change_pct"`
	MarketCapChange    string `json:"market_cap_change"`
	MarketCapChangePct string `json:"market_cap_change_pct"`
}
