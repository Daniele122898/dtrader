package api

import (
	"dtrader/pkg/gnomics/internal/bc"
	"dtrader/pkg/gnomics/internal/qp"
	"dtrader/pkg/gnomics/internal/ub"
	"dtrader/pkg/gnomics/models"
	"encoding/json"
	"errors"
	"time"
)

// Get current exchange rates for prices from markets into USD. Also includes BTC and ETH.
// For example it will have an entry of BTC with a value of 48000. Thus 1 BTC is worth 48k USD.
func (g *Gnomics) GetExchangeRates() ([]models.ExchangeRate, error) {
	params := make (qp.QueryParams, 1)
	resp, err := g.getRequest(
		ub.BuildUrlSt("exchange-rates"),
		params)
	if err != nil {
		return nil, err
	}

	l := bc.CountRune(resp, '{')
	data := make([]models.ExchangeRate, 0, l)
	err = json.Unmarshal(resp, &data)

	if err != nil {
		return nil, err
	}

	return data, err
}

// Get Volume history of all crypto assets between time period
// The currency parameter must be a Nomics Quote Currency, to get all Nomics Quote Currencies, \
// use the /exchange-rates endpoint for all current rates.
//	currency (required): The currency we want the exchange rate from
//	start (required): Set start period.
//	end: End date of period. If not set today is used
func (g *Gnomics) GetExchangeRateHistory(currency string, start time.Time, end time.Time) ([]models.ExchangeRateHistory, error) {
	params := make (qp.QueryParams, 4)

	if currency == "" {
		return nil, errors.New("currency must be set")
	}
	params["currency"] = currency
	if start.IsZero() {
		return nil, errors.New("start time must be set")
	}
	params["start"] = start.Format(time.RFC3339)
	if !end.IsZero() {
		params["end"] = end.Format(time.RFC3339)
	}

	resp, err := g.getRequest(
		ub.BuildUrlSt("exchange-rates/history"),
		params)

	if err != nil {
		return nil, err
	}

	count := bc.CountRune(resp, '{')
	data := make([]models.ExchangeRateHistory, 0, count)
	err = json.Unmarshal(resp, &data)

	if err != nil {
		return nil, err
	}

	return data, err
}
