package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
	"work-project/internal/airtable"
)

type AirTable interface {
	GetProducts(ctx context.Context) ([]airtable.BaseObject[airtable.ProductListResponse], error)
}

type AirTableClient struct {
	client  *http.Client
	baseURL *url.URL
	apiKey  string
}

func NewAirTableClient(baseUrl string, apiKey string) (*AirTableClient, error) {
	baseURL, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	return &AirTableClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: baseURL,
		apiKey:  apiKey,
	}, nil
}

func (r *AirTableClient) GetProducts(ctx context.Context) ([]airtable.BaseObject[airtable.ProductListResponse], error) {
	requestURL := r.baseURL.JoinPath("/Store")
	req, err := r.newRequest(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error while getting Airtable products. Response code: %d", resp.StatusCode)
	}

	rawResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response airtable.BaseResponse[airtable.ProductListResponse]
	if err := json.Unmarshal(rawResponse, &response); err != nil {
		return nil, err
	}

	return response.Records, nil
}

func (r *AirTableClient) newRequest(
	ctx context.Context,
	method string,
	requestURL *url.URL,
	body interface{},
) (*http.Request, error) {
	var bodyReader io.Reader
	if body != nil {
		rawBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(rawBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, requestURL.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	req.Header.Add("Authorization", "Bearer "+r.apiKey)

	return req, nil
}
