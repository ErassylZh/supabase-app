package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
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
			Timeout: 30,
		},
		baseURL: baseURL,
		apiKey:  apiKey,
	}, nil
}

func (r *AirTableClient) GetProducts(ctx context.Context) ([]airtable.BaseObject[airtable.ProductListResponse], error) {
	requestURL := r.baseURL.JoinPath("/Stores")
	req, err := r.newRequest(ctx, http.MethodPost, requestURL, nil)
	if err != nil {
		return []airtable.BaseObject[airtable.ProductListResponse]{}, err
	}
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error while get airtable products. Response code " + strconv.Itoa(resp.StatusCode))
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
	requestUrl *url.URL,
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

	req, err := http.NewRequestWithContext(ctx, method, requestUrl.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	authHeader := fmt.Sprintf("%s %s", "Bearer", r.apiKey)
	if authHeader != "" {
		req.Header.Add("Authorization", authHeader)
	}

	return req, nil
}
