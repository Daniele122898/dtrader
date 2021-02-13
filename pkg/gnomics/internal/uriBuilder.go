package gnomics

import (
	"dtrader/pkg/gnomics"
	"fmt"
	"strings"
)

// buildUriSt will add upon the default endpoint and create a full request
// url. This method uses strings only and thus does not do any conversions.
// For automatic conversions use the general buildUrl function
func buildUrlSt(parts ...string) string {
	var sb strings.Builder
	sb.WriteString(gnomics.ENDPOINT)

	for _, p := range parts {
		sb.WriteRune('/')
		sb.WriteString(p)
	}

	return sb.String()
}

// buildUrl will add upon the default endpoint and create a full request URL.
// This method will attempt conversions and thus is slightly slower.
func buildUrl(parts ...interface{}) string {
	var sb strings.Builder
	sb.WriteString(gnomics.ENDPOINT)

	for _, p := range parts {
		sb.WriteRune('/')
		sb.WriteString(fmt.Sprint(p))
	}

	return sb.String()
}