package api

import "dtrader/pkg/gnomics/models"

const (
	Interval1H = "1h"
	Interval1D = "1d"
	Interval7D = "7d"
	Interval30D = "30d"
	Interval356d = "365d"
	IntervalYtd = "ytd"
)

// Fetches Currency data like volume, price etc.
//	ids: The internal IDs of currencies that we want to fetch. If nil is passed it will fetch all currencies available
//	(be aware that the return size is in the orders of mb)
//	interval: The intervals to request. If nil is passed it will request 1d-ytd
//	convert: To which currency to convert to. Default is USD
func (g *Gnomics) GetCurrenciesTicker(ids []string, interval []string, convert *string) ([]models.CurrencyTicker, error) {

}