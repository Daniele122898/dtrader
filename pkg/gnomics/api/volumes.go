package api

import (
	"dtrader/pkg/gnomics/internal/bc"
	"dtrader/pkg/gnomics/internal/qp"
	"dtrader/pkg/gnomics/internal/ub"
	"dtrader/pkg/gnomics/models"
	"encoding/json"
	"time"
)

// Get Volume history of all crypto assets between time period
//	start: Set start period. If not set it just fetches form the first possible data point
//	end: End date of period. If not set today is used
//	convert: To what currency to convert to. Defaults to USD
func (g *Gnomics) GetVolume(start time.Time, end time.Time, convert string) ([]models.Volume, error) {
	params := make (qp.QueryParams, 4)

	if !start.IsZero() {
		params["start"] = start.Format(time.RFC3339)
	}
	if !end.IsZero() {
		params["end"] = end.Format(time.RFC3339)
	}
	if convert != "" {
		params["convert"] = convert
	}

	resp, err := g.getRequest(
		ub.BuildUrlSt("volume/history"),
		params)

	if err != nil {
		return nil, err
	}

	count := bc.CountRune(resp, '{')
	data := make([]models.Volume, 0, count)
	err = json.Unmarshal(resp, &data)

	if err != nil {
		return nil, err
	}

	return data, err
}
