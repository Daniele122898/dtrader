package api

import (
	"dtrader/pkg/gnomics/internal/qp"
	"errors"
	"net/http"
	"strings"
)

type Gnomics struct {
	apiKey string
	httpClient http.Client
}

func NewGnomics(apiKey string) (*Gnomics, error) {
	if len(apiKey) == 0 {
		return nil, errors.New("you must specify an API key")
	}

	if strings.HasPrefix(apiKey, "demo") {
		return nil, errors.New("do not use demo api keys but request your own")
	}

	// THIS IS VERY IMPORTANT. Otherwise all receiver funcs will always copy the val.
	// It must be a reference from the start
	return &Gnomics{
		apiKey: apiKey,
		httpClient: http.Client{},
	}, nil
}

func (g *Gnomics) injectApiKey(q qp.QueryParams) error {
	if q == nil {
		return errors.New("query params was not initialized")
	}
	if _, ok := q["key"]; !ok {
		q["key"] = g.apiKey
	}
	return nil
}
