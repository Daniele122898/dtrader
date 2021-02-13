package test

import (
	"dtrader/pkg/gnomics"
	"testing"
)

func TestGetCurrenciesTicker(t *testing.T) {
	c, err := gnomics.NewGnomics(demoApiKey)
	if err != nil {
		t.Errorf("couldn't create client %v", err)
	}

	r, err := c.GetCurrenciesTicker([]string{"BTC"}, []string{gnomics.Interval1H}, nil)
	if err != nil {
		t.Errorf("couldn't create client %v", err)
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

	if r[0].ID != "BTC" {
		t.Errorf("want: BTC, got %v", r[0].ID)
		return
	}
}