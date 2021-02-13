package api

import (
	"dtrader/pkg/gnomics/internal/qp"
	"dtrader/pkg/gnomics/internal/ub"
	"dtrader/pkg/gnomics/models"
	"strings"
)

// Fetches Currency data like volume, price etc.
//	ids: The internal IDs of currencies that we want to fetch. If nil is passed it will fetch all currencies available
//	(be aware that the return size is in the orders of mb)
//	interval: The intervals to request. If nil is passed it will request 1d-ytd
//	convert: To which currency to convert to. Default is USD
func (g *Gnomics) GetCurrenciesTicker(ids []string, interval []string, convert *string) ([]models.CurrencyTicker, error) {
	params := make (qp.QueryParams, 4)
	l := 12000 // In case ALL currencies are requested
	if ids != nil {
		params["ids"] = strings.Join(ids, ",")
		l = len(ids)
	}
	if interval != nil {
		params["interval"] = strings.Join(interval, ",")
	}
	if convert != nil {
		params["convert"] = *convert
	}

	data := make([]models.CurrencyTicker, l)
	err := g.getRequestParsed(
		ub.BuildUrlSt("currencies/ticker"),
		params,
		&data)

	if err != nil {
		return nil, err
	}

	return data, err
}

// Fetches Currency data like volume, price etc.
//	ids: The internal IDs of currencies that we want to fetch. If nil is passed it will fetch all currencies available
//	(be aware that the return size is in the orders of mb)
//	attributes: The attributes to request. If nil is passed it will request 1d-ytd
func (g *Gnomics) GetCurrenciesMetadata(ids []string, attributes []string) ([]models.CurrencyMetadata, error) {
	params := make (qp.QueryParams, 3)
	l := 12000 // In case ALL currencies are requested
	if ids != nil {
		params["ids"] = strings.Join(ids, ",")
		l = len(ids)
	}
	if attributes != nil {
		params["attributes"] = strings.Join(attributes, ",")
	}

	data := make([]models.CurrencyMetadata, l)
	err := g.getRequestParsed(
		ub.BuildUrlSt("currencies"),
		params,
		&data)

	if err != nil {
		return nil, err
	}

	return data, err
}