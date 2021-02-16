package gnomics

import (
	"dtrader/pkg/gnomics/api"
)

type gnomics = api.Gnomics

func NewGnomics(apiKey string) (*gnomics, error) {
	return api.NewGnomics(apiKey)
}
