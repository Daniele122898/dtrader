package api

import (
	"dtrader/pkg/gnomics/internal/bc"
	"dtrader/pkg/gnomics/internal/qp"
	"dtrader/pkg/gnomics/internal/ub"
	"dtrader/pkg/gnomics/models"
	"encoding/json"
	"strings"
)

// Fetches Market information about which exchanges provide the exchange between base and quote
//	exchange: Which exchange we are interested in. If left blank it will fetch ALL exchanges. Prepare for massive transfers
//	base: Which base currencies you are interested in. If left blank it will show all base permutations of base currencies
//	available in the exchange for the set quote
//	quote: The quote for the base currency. If not set it will show all available
func (g *Gnomics) GetMarkets(exchange string, base []string, quote []string) ([]models.Market, error)  {
	params := make(qp.QueryParams, 4)
	if exchange != "" {
		params["exchange"] = exchange
	}
	if base != nil {
		params["base"] = strings.Join(base, ",")
	}
	if quote != nil {
		params["quote"] = strings.Join(quote, ",")
	}

	resp, err := g.getRequest(
		ub.BuildUrlSt("markets"),
		params)

	count := bc.CountRune(resp, '{')
	data := make([]models.Market, 0, count)
	err = json.Unmarshal(resp, &data)

	if err != nil {
		return nil, err
	}

	return data, err

}