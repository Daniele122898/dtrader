package api

import (
	"dtrader/pkg/gnomics/internal/bc"
	"dtrader/pkg/gnomics/internal/qp"
	"dtrader/pkg/gnomics/internal/ub"
	"dtrader/pkg/gnomics/models"
	"encoding/json"
	"testing"
	"time"
)

func TestGetMarkets(t *testing.T) {
	c, err := NewGnomics(demoApiKey)
	if err != nil {
		t.Errorf("couldn't create client %v", err)
	}

	r, err := c.GetMarkets("binance", []string{"BTC"}, nil)

	if err != nil {
		t.Errorf("couldn't get a response %v", err)
		return
	}

	if r == nil {
		t.Errorf("response is nil")
		return
	}

	if len(r) == 0 {
		t.Errorf("response array is empty")
		return
	}

	if r[0].Base != "BTC" {
		t.Errorf("want: BTC, got %v", r[0].Base)
		return
	}
}

func BenchmarkTestGetMarketsAll(b *testing.B) {
	g, err := NewGnomics(demoApiKey)
	if err != nil {
		b.Errorf("couldn't create client %v", err)
	}

	params := make (qp.QueryParams, 1)

	resp, err := g.getRequest(
		ub.BuildUrlSt("markets"),
		params)
	if err != nil {
		b.Errorf("couldn't make request %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l := bc.CountRune(resp, '{')

		data := make([]models.Market, 0, l)
		err = json.Unmarshal(resp, &data)
		if err != nil {
			b.Errorf("couldn't marshall %v", err)
		}
	}
}

func TestGetMarketCap(t *testing.T) {
	c, err := NewGnomics(demoApiKey)
	if err != nil {
		t.Errorf("couldn't create client %v", err)
	}

	start := time.Now().AddDate(0,0,-10)
	r, err := c.GetMarketCap(start, time.Time{}, "USD")

	if err != nil {
		t.Errorf("couldn't get a response: %v", err)
		return
	}

	if r == nil {
		t.Errorf("response is nil")
		return
	}

	if len(r) != 10 {
		t.Errorf("want: length 10, got: %v", len(r))
		return
	}

	if len(r[0].MarketCap) < 3 {
		t.Errorf("want: market cap that is longer than 3, got %v", len(r[0].MarketCap))
		return
	}
}
