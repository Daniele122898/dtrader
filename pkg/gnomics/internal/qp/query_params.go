package qp

import (
	"net/url"
	"strings"
)

type QueryParams map[string]string

func (q QueryParams) parse() string {
	if len(q) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString("?")
	for param, val := range q {
		sb.WriteString(url.QueryEscape(param))
		sb.WriteRune('=')
		sb.WriteString(url.QueryEscape(val))
		sb.WriteRune('&')
	}

	return sb.String()[:sb.Len() - 1]
}
