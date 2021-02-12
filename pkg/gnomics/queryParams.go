package gnomics

import "strings"

type QueryParams map[string]string

func (q QueryParams) parse() string {
	if len(q) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString("?")
	for param, val := range q {
		sb.WriteString(param)
		sb.WriteRune('=')
		sb.WriteString(val)
		sb.WriteRune('&')
	}

	return sb.String()[:sb.Len() - 1]
}
