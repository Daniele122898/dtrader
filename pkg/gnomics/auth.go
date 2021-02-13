package gnomics

import (
	"errors"
	"strings"
)

var (
	apiKey = ""
)

func HasApiKey() bool {
	return len(apiKey) > 0
}

func Authenticate(key string) error {
	if len(key) == 0 {
		return errors.New("you must specify an API key")
	}

	if strings.HasPrefix(key, "demo") {
		return errors.New("do not use demo api keys but request your own")
	}

	apiKey = key
	return nil
}

func injectApiKey(q QueryParams) error {
	if q == nil {
		return errors.New("query params was not initialized")
	}
	if _, ok := q["key"]; !ok {
		q["key"] = apiKey
	}
	return nil
}
