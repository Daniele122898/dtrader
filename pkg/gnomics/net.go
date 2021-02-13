package gnomics

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var client = http.Client{}

func prepareQueryForAuth(q *QueryParams) error {
	if *q == nil {
		*q = make(map[string]string)
	}
	if err := InjectApiKey(*q); err != nil {
		return err
	}
	return nil
}

// Will execute a get-request, inject the api key if it exists and then parse the data.
func getRequestParsed(endpoint string, queryParams QueryParams, data *interface{}) error {
	if err := prepareQueryForAuth(&queryParams); err != nil {
		return err
	}

	resp, err := request(endpoint, http.MethodGet, queryParams, nil)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, data)
	return err
}

// Will execute a getRequest and inject the api key if it exists
func getRequest(endpoint string, queryParams QueryParams) ([]byte, error) {
	if err := prepareQueryForAuth(&queryParams); err != nil {
		return nil, err
	}

	return request(endpoint, http.MethodGet, queryParams, nil)
}

// Will execute a postRequest and inject the api key if it exists
func postRequest(endpoint string, queryParams QueryParams, body interface{}) ([]byte, error) {
	if err := prepareQueryForAuth(&queryParams); err != nil {
		return nil, err
	}

	return request(endpoint, http.MethodPost, queryParams, body)
}

// General request function. This just executes and passes the errors or data.
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