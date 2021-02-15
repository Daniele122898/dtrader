package test

import (
	"dtrader/pkg/gnomics"
	"testing"
	"time"
)

func TestGetCurrenciesTicker(t *testing.T) {
	c, err := gnomics.NewGnomics(demoApiKey)
	if err != nil {
		t.Errorf("couldn't create client %v", err)
	}

	r, err := c.GetCurrenciesTicker([]string{"BTC"}, []string{gnomics.Interval1H}, nil)
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

	if r[0].ID != "BTC" {
		t.Errorf("want: BTC, got %v", r[0].ID)
		return
	}
}

func BenchmarkGetCurrenciesTicker(b *testing.B) {
	c, err := gnomics.NewGnomics(demoApiKey)
	if err != nil {
		b.Errorf("couldn't create client %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.GetCurrenciesTicker([]string{"BTC"}, []string{gnomics.Interval1H}, "")
	}
}

func BenchmarkGetCurrenciesTickerAll(b *testing.B) {
	c, err := gnomics.NewGnomics(demoApiKey)
	if err != nil {
		b.Errorf("couldn't create client %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.GetCurrenciesTicker(nil, []string{gnomics.Interval1H}, "")
	}
}

func TestGetCurrenciesMetadata(t *testing.T) {
	c, err := gnomics.NewGnomics(demoApiKey)
	if err != nil {
		t.Errorf("couldn't create client %v", err)
	}

	r, err := c.GetCurrenciesMetadata([]string{"BTC"}, nil)
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

	if r[0].ID != "BTC" {
		t.Errorf("want: BTC, got %v", r[0].ID)
		return
	}
}

func BenchmarkGetCurrenciesMetadata(b *testing.B) {
	c, err := gnomics.NewGnomics(demoApiKey)
	if err != nil {
		b.Errorf("couldn't create client %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.GetCurrenciesMetadata([]string{"BTC"}, nil)
	}
}

func BenchmarkGetCurrenciesMetadataAll(b *testing.B) {
	c, err := gnomics.NewGnomics(demoApiKey)
	if err != nil {
		b.Errorf("couldn't create client %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.GetCurrenciesMetadata(nil, nil)
	}
}

func TestGetCurrenciesSparkline(t *testing.T) {
	c, err := gnomics.NewGnomics(demoApiKey)
	if err != nil {
		t.Errorf("couldn't create client %v", err)
	}

	start := time.Now().AddDate(0,0,-10)
	r, err := c.GetCurrenciesSparkline([]string{"BTC"}, start, time.Time{}, "")
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

	p := r[0]

	if len(p.Timestamps) != 10 {
		t.Errorf("want: 10 timestamps, got: %v", len(p.Timestamps))
		return
	}

	if len(p.Prices) != 10 {
		t.Errorf("want: 10 prices, got: %v", len(p.Prices))
		return
	}

	if p.Currency != "BTC" {
		t.Errorf("want: BTC, got %v", p.Currency)
		return
	}
}

func BenchmarkGetCurrenciesSparkline45(b *testing.B) {
	c, err := gnomics.NewGnomics(demoApiKey)
	if err != nil {
		b.Errorf("couldn't create client %v", err)
	}
	start := time.Now().AddDate(0,0,-45)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.GetCurrenciesSparkline([]string{"BTC"}, start, time.Time{}, "")
	}
}

func BenchmarkGetCurrenciesSparkline45All(b *testing.B) {
	c, err := gnomics.NewGnomics(demoApiKey)
	if err != nil {
		b.Errorf("couldn't create client %v", err)
	}
	start := time.Now().AddDate(0,0,-45)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.GetCurrenciesSparkline(nil, start, time.Time{}, "")
	}
}