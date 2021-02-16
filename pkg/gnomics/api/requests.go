package api

import (
	"bytes"
	"dtrader/pkg/gnomics/internal/qp"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Will execute a get-request, inject the api key if it exists and then parse the data.
func (g *Gnomics) getRequestParsed(endpoint string, queryParams qp.QueryParams, data interface{}) error {
	resp, err := g.request(endpoint, http.MethodGet, queryParams, nil)
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp, data)
	return err
}

// Will execute a getRequest and inject the api key if it exists
func (g *Gnomics) getRequest(endpoint string, queryParams qp.QueryParams) ([]byte, error) {
	return g.request(endpoint, http.MethodGet, queryParams, nil)
}

// Will execute a postRequest and inject the api key if it exists
func (g *Gnomics) postRequest(endpoint string, queryParams qp.QueryParams, body interface{}) ([]byte, error) {
	return g.request(endpoint, http.MethodPost, queryParams, body)
}

func prepareQueryForAuth(q *qp.QueryParams, g *Gnomics) error {
	if *q == nil {
		*q = make(map[string]string)
	}
	if err := g.injectApiKey(*q); err != nil {
		return err
	}
	return nil
}

// General request function. This just executes and passes the errors or data.
func (g *Gnomics) request(endpoint string, httpMethod string, queryParams qp.QueryParams, body interface{}) ([]byte, error) {
	if err := prepareQueryForAuth(&queryParams, g); err != nil {
		return nil, err
	}

	url := endpoint
	if queryParams != nil {
		url += queryParams.Parse()
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

	resp, err := g.httpClient.Do(req)
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