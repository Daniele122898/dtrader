package gnomics

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var client = http.Client{}

func getRequest(endpoint string, queryParams QueryParams) ([]byte, error) {
	return request(endpoint, http.MethodGet, queryParams, nil)
}

func postRequest(endpoint string, queryParams QueryParams, body interface{}) ([]byte, error) {
	return request(endpoint, http.MethodPost, queryParams, body)
}

func request(endpoint string, httpMethod string, queryParams QueryParams, body interface{}) ([]byte, error) {
	url := endpoint
	if queryParams != nil {
		url += queryParams.parse()
	}

	var req *http.Request
	var err error

	switch httpMethod {
	case http.MethodGet:
		req, err = http.NewRequest(httpMethod, url, nil)
	case http.MethodPost:
		js, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(httpMethod, url, bytes.NewBuffer(js))
	default:
		return nil, errors.New("unsupported HTTP method")
	}
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// Close the body so we dont leak any mem
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}