package api

import (
	"testing"
)

func TestGetExchangeRates(t *testing.T) {
	c, err := NewGnomics(demoApiKey)
	if err != nil {
		t.Errorf("couldn't create client %v", err)
	}

	r, err := c.GetExchangeRates()

	if err != nil {
		t.Errorf("couldn't get a response: %v", err)
		return
	}

	if r == nil {
		t.Errorf("response is nil")
		return
	}

	if len(r) == 0 {
		t.Errorf("want: length bigger than 0, got: %v", len(r))
		return
	}

	if r[0].Rate == "" {
		t.Errorf("want: rate to be non empty string, got %v", r[0].Rate)
		return
	}
}
