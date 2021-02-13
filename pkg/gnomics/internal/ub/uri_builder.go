package ub

import (
	"fmt"
	"strings"
)

const (
	ENDPOINT = "https://api.nomics.com/v1"
)

// BuildUrlSt will add upon the default endpoint and create a full request
// url. This method uses strings only and thus does not do any conversions.
// For automatic conversions use the general BuildUrl function
func BuildUrlSt(parts ...string) string {
	var sb strings.Builder
	sb.WriteString(ENDPOINT)

	for _, p := range parts {
		sb.WriteRune('/')
		sb.WriteString(p)
	}

	return sb.String()
}

// BuildUrl will add upon the default endpoint and create a full request URL.
// This method will attempt conversions and thus is slightly slower.
func BuildUrl(parts ...interface{}) string {
	var sb strings.Builder
	sb.WriteString(ENDPOINT)

	for _, p := range parts {
		sb.WriteRune('/')
		sb.WriteString(fmt.Sprint(p))
	}

	return sb.String()
}