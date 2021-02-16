package api

import (
	"dtrader/pkg/gnomics/internal/bc"
	"dtrader/pkg/gnomics/internal/qp"
	"dtrader/pkg/gnomics/internal/ub"
	"dtrader/pkg/gnomics/models"
	"encoding/json"
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
