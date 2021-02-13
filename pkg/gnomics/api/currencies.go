package api

import (
	"dtrader/pkg/gnomics/internal/qp"
	"dtrader/pkg/gnomics/internal/ub"
	"dtrader/pkg/gnomics/models"
	"errors"
	"strings"
	"time"
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

// Fetches Currency data like metadata like description, urls etc.
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

// Fetches Currency Sparkline (Prices)
//	ids: The internal IDs of currencies that we want to fetch. If nil is passed it will fetch all currencies available
//	(be aware that the return size is in the orders of mb)
//	start: start time of the interval
//	end: end time of the interval, if not set will choose current time and date
//	convert: To which currency to convert to. Default is USD
//	!! NOTE: To get information on a daily level start and end can be a maximum of 45 days apart. Otherwise it wills tart to accumulate the prices and show only specific days.
func (g *Gnomics) GetCurrenciesSparkline(ids []string, start *time.Time, end *time.Time, convert *string) ([]models.CurrencySparkline, error) {
	params := make (qp.QueryParams, 3)
	l := 12000 // In case ALL currencies are requested
	if ids != nil {
		params["ids"] = strings.Join(ids, ",")
		l = len(ids)
	}
	if start == nil {
		return nil, errors.New("start time must be set")
	}
	params["start"] = start.Format(time.RFC3339)
	if end != nil {
		params["end"] = end.Format(time.RFC3339)
	}
	if convert != nil {
		params["convert"] = *convert
	}

	data := make([]models.CurrencySparkline, l)

	err := g.getRequestParsed(
		ub.BuildUrlSt("currencies/sparkline"),
		params,
		&data)

	if err != nil {
		return nil, err
	}

	return data, err
}