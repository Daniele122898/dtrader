package test

import (
	"dtrader/pkg/gnomics"
	"testing"
)

func TestGetMarkets(t *testing.T) {
	c, err := gnomics.NewGnomics(demoApiKey)
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

func BenchmarkTestGetMarkets(b *testing.B) {
	c, err := gnomics.NewGnomics(demoApiKey)
	if err != nil {
		b.Errorf("couldn't create client %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.GetMarkets("binance", []string{"BTC"}, nil)
	}
}

func BenchmarkTestGetMarketsAll(b *testing.B) {
	c, err := gnomics.NewGnomics(demoApiKey)
	if err != nil {
		b.Errorf("couldn't create client %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.GetMarkets("", nil, nil)
	}
}
